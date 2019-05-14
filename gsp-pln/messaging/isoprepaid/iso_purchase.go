package isoprepaid

import (
	"encoding/json"
	"log"

	"github.com/Ayvan/iso8583"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/config"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/basic"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/util"
)

// IsoPurchase : iso purchase for prepaid
type IsoPurchase struct {
}

// Encode : to encode message for prepaid purchase iso8583 format
func (isoPurchase *IsoPurchase) Encode(msgJSON string) []byte {

	log.Println("prepaid.IsoPurchase[Encode(message string)] : start to encode")
	message := &basic.Message{
		AdditionalPrivateData:  &AdditionalPrivateData{},
		AdditionalPrivateData3: &AdditionalPrivateData3{},
	}

	log.Println("prepaid.IsoPurchase[Encode(message string)] : encode json format to iso")
	isoFormat, msgPurchase := basic.EncodeJSONFormatToISO(msgJSON, message)

	isoFormat.TransactionAmount = iso8583.NewAlphanumeric(basic.FormatTrxAmountString(msgPurchase.TransactionAmount))
	isoFormat.AdditionalPrivateData =
		iso8583.NewLllvar([]byte(FormatPurchaseReq(msgPurchase.AdditionalPrivateData.(*AdditionalPrivateData))))

	isoFormat.AdditionalPrivateData3 =
		iso8583.NewLllvar([]byte(FormatData3String(msgPurchase.AdditionalPrivateData3.(*AdditionalPrivateData3))))
	if msgPurchase.Mti == config.Get().Mti.Payment.Response {
		isoFormat.ResponseCode = iso8583.NewAlphanumeric(msgPurchase.ResponseCode)
		if msgPurchase.ResponseCode == "0000" {
			isoFormat.SettlementDate = iso8583.NewAlphanumeric(msgPurchase.SettlementDate)
			isoFormat.AdditionalPrivateData =
				iso8583.NewLllvar([]byte(FormatPurchaseRes(msgPurchase.AdditionalPrivateData.(*AdditionalPrivateData))))
			isoFormat.AdditionalPrivateData2 =
				iso8583.NewLllvar([]byte(msgPurchase.AdditionalPrivateData2.(string)))

		}
	}

	msg := iso8583.NewMessageExtended(msgPurchase.Mti, iso8583.ASCII, false, true, isoFormat)

	packetIso, err := msg.Bytes()
	if err != nil {
		panic(err.Error())
	}

	log.Println("prepaid.IsoPurchase[Encode(message string)] : end to encode")
	return util.EncapsulateBytes(packetIso)

}

// Decode : decode from byte iso8583 to prepaid purchase
func (isoPurchase *IsoPurchase) Decode(message []byte) (string, error) {

	log.Println("prepaid.IsoPurchase[Decode(message string)] : start to decode")
	resultFields, mti := basic.DecodeIsoMessage(message)

	log.Println("prepaid.IsoPurchase[Decode(message string)] : start to assign iso to message")
	msgPurResult := basic.AssignISOFormatToMessage(resultFields, mti)

	if mti == config.Get().Mti.Payment.Response || mti == config.Get().Mti.Advice.Response || mti == config.Get().Mti.Advice.Repeat.Response {
		msgPurResult.ResponseCode = resultFields.ResponseCode.Value
		msgPurResult.TransactionAmount = basic.ParseMessageToTrxAmt(resultFields.TransactionAmount.Value)

		if resultFields.ResponseCode.Value == "0000" {
			msgPurResult.SettlementDate = resultFields.SettlementDate.Value
			msgPurResult.AdditionalPrivateData = BuildPurchaseResponse(string(resultFields.AdditionalPrivateData.Value))
			msgPurResult.AdditionalPrivateData2 = string(resultFields.AdditionalPrivateData2.Value)
			msgPurResult.AdditionalPrivateData3 = BuildData3Response(string(resultFields.AdditionalPrivateData3.Value))
		} else {
			msgPurResult.AdditionalPrivateData = BuildInquiryResponse(string(resultFields.AdditionalPrivateData.Value))
		}
	}

	log.Println("prepaid.IsoPurchase[Decode(message string)] json decode : ", &msgPurResult)

	json, _ := json.Marshal(msgPurResult)

	return string(json), nil

}
