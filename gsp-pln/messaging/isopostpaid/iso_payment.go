package isopostpaid

import (
	"encoding/json"
	"log"

	"github.com/Ayvan/iso8583"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/config"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/basic"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/util"
)

// IsoPayment : iso payment postpaid
type IsoPayment struct {
}

// Encode : to encode message for postpaid payment iso8583 format
func (isoPayment *IsoPayment) Encode(msgJSON string) []byte {

	log.Println("postpaid.IsoPayment[Encode(message string)] : start to encode")
	message := &basic.Message{
		AdditionalPrivateData:  &AdditionalPrivateData{},
		AdditionalPrivateData2: &AdditionalPrivateData2{},
	}

	log.Println("postpaid.IsoInquiry[Encode(message string)] : encode json format to iso")
	isoFormat, msgPayment := basic.EncodeJSONFormatToISO(msgJSON, message)

	isoFormat.AdditionalPrivateData =
		iso8583.NewLllvar([]byte(FormatDataString(msgPayment.AdditionalPrivateData.(*AdditionalPrivateData))))

	if msgPayment.Mti == config.Get().Mti.Payment.Response {
		isoFormat.ResponseCode = iso8583.NewAlphanumeric(msgPayment.ResponseCode)
		if msgPayment.ResponseCode == "0000" {
			isoFormat.SettlementDate = iso8583.NewAlphanumeric(msgPayment.SettlementDate)
			isoFormat.AdditionalPrivateData2 =
				iso8583.NewLllvar([]byte(FormatData2String(msgPayment.AdditionalPrivateData2.(*AdditionalPrivateData2))))
		}
	}

	msg := iso8583.NewMessageExtended(msgPayment.Mti, iso8583.ASCII, false, true, isoFormat)

	packetIso, err := msg.Bytes()
	if err != nil {
		panic(err.Error())
	}

	log.Println("postpaid.IsoPayment[Encode(message string)] : end to encode")
	return util.EncapsulateBytes(packetIso)

}

// Decode : decode from byte iso8583 to postpaid payment
func (isoPayment *IsoPayment) Decode(message []byte) (string, error) {

	log.Println("postpaid.IsoPayment[Decode(message string)] : start to decode")
	resultFields, mti := basic.DecodeIsoMessage(message)

	log.Println("postpaid.IsoInquiry[Decode(message string)] : start to assign iso to message")
	msgPayResult := basic.AssignISOFormatToMessage(resultFields, mti)

	msgPayResult.AdditionalPrivateData = BuildDataResponse(string(resultFields.AdditionalPrivateData.Value))

	if mti == config.Get().Mti.Payment.Response {
		msgPayResult.ResponseCode = resultFields.ResponseCode.Value
		if resultFields.ResponseCode.Value == "0000" {
			msgPayResult.SettlementDate = resultFields.SettlementDate.Value
			msgPayResult.AdditionalPrivateData2 = BuildData2Response(string(resultFields.AdditionalPrivateData2.Value))
		}
	}

	log.Println("postpaid.IsoPayment[Decode(message string)] json decode : ", &msgPayResult)

	json, _ := json.Marshal(msgPayResult)

	return string(json), nil

}
