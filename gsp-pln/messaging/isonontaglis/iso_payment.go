package isonontaglis

import (
	"encoding/json"

	log "gitlab.com/kasku/kasku-2pay/2pay-billerpayment/log"

	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/config"

	"github.com/Ayvan/iso8583"

	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/basic"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/util"
)

// IsoPayment : non tagihan listrik payment
type IsoPayment struct {
}

// Encode : to encode message for nontaglis payment
func (isoPayment *IsoPayment) Encode(msgJSON string) []byte {

	log.Get().Println("nontaglis.IsoPayment[Encode(message string)] : start to encode")
	message := &basic.Message{
		AdditionalPrivateData:  &AdditionalPrivateData{},
		AdditionalPrivateData2: &AdditionalPrivateData2{},
		AdditionalPrivateData3: &AdditionalPrivateData3{},
	}

	log.Get().Println("nontaglis.IsoInquiry[Encode(message string)] : encode json format to iso")
	isoFormat, msgPayment := basic.EncodeJSONFormatToISO(msgJSON, message)

	isoFormat.TransactionAmount = iso8583.NewAlphanumeric(basic.FormatTrxAmountString(msgPayment.TransactionAmount))

	isoFormat.AdditionalPrivateData = iso8583.NewLllvar([]byte(FormatDataString(msgPayment.AdditionalPrivateData.(*AdditionalPrivateData))))
	isoFormat.AdditionalPrivateData3 = iso8583.NewLllvar([]byte(FormatData3String(msgPayment.AdditionalPrivateData3.(*AdditionalPrivateData3))))

	if msgPayment.Mti == config.Get().Mti.Payment.Response {
		isoFormat.ResponseCode = iso8583.NewAlphanumeric(msgPayment.ResponseCode)
		if msgPayment.ResponseCode == "0000" {
			isoFormat.SettlementDate = iso8583.NewAlphanumeric(msgPayment.SettlementDate)
			isoFormat.AdditionalPrivateData2 = iso8583.NewLllvar([]byte(FormatData2String(msgPayment.AdditionalPrivateData2.(*AdditionalPrivateData2))))
			isoFormat.InfoText = iso8583.NewLllvar([]byte(msgPayment.InfoText))
		}
	}

	msg := iso8583.NewMessageExtended(msgPayment.Mti, iso8583.ASCII, false, true, isoFormat)

	packetIso, err := msg.Bytes()
	if err != nil {
		panic(err.Error())
	}

	return util.EncapsulateBytes(packetIso)
}

// Decode : decode from byte iso8583 to nontaglis payment
func (isoPayment *IsoPayment) Decode(message []byte) (string, error) {

	log.Get().Println("nontaglis.IsoPayment[Decode(message string)] : start to decode")
	resultFields, mti := basic.DecodeIsoMessage(message)

	log.Get().Println("nontaglis.IsoInquiry[Decode(message string)] : start to assign iso to message")
	msgPayResult := basic.AssignISOFormatToMessage(resultFields, mti)

	msgPayResult.TransactionAmount = basic.ParseMessageToTrxAmt(resultFields.TransactionAmount.Value)
	msgPayResult.AdditionalPrivateData = BuildResponse(string(resultFields.AdditionalPrivateData.Value))
	msgPayResult.AdditionalPrivateData3 = BuildData3Respose(string(resultFields.AdditionalPrivateData3.Value))

	if mti == config.Get().Mti.Payment.Response {
		msgPayResult.ResponseCode = resultFields.ResponseCode.Value
		if resultFields.ResponseCode.Value == "0000" {
			msgPayResult.SettlementDate = resultFields.SettlementDate.Value
			msgPayResult.AdditionalPrivateData2 = BuildData2Response(string(resultFields.AdditionalPrivateData2.Value))
			msgPayResult.InfoText = string(resultFields.InfoText.Value)
		}
	}

	log.Get().Println("nontaglis.IsoPayment[Decode(message string)] json decode : ", &msgPayResult)

	json, _ := json.Marshal(msgPayResult)

	return string(json), nil
}
