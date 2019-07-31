package isoprepaid

import (
	"encoding/json"
	"log"

	"github.com/Ayvan/iso8583"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/config"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/basic"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/util"
)

// IsoInquiry : iso inquiry for prepaid
type IsoInquiry struct {
}

// Encode : to encode message for prepaid inquiry
func (isoInquiry *IsoInquiry) Encode(msgJSON string) []byte {
	log.Println("prepaid.IsoInquiry[Encode(message string)] : start to encode ")

	log.Println("prepaid.IsoInquiry[Encode(message string)] : initialize message to assign interface with isoprepaid message")
	message := &basic.Message{
		AdditionalPrivateData:  &AdditionalPrivateData{},
		AdditionalPrivateData3: &AdditionalPrivateData3{},
	}

	log.Println("prepaid.IsoInquiry[Encode(message string)] : encode json format to iso")
	isoFormat, msgInquiry := basic.EncodeJSONFormatToISO(msgJSON, message)

	if msgInquiry.Mti == config.Get().Mti.Inquiry.Request {
		isoFormat.AdditionalPrivateData =
			iso8583.NewLllvar([]byte(FormatInqReq(msgInquiry.AdditionalPrivateData.(*AdditionalPrivateData))))
	} else if msgInquiry.Mti == config.Get().Mti.Inquiry.Response {
		if len(msgInquiry.TransactionAmount.ValueAmount) > 0 {
			isoFormat.TransactionAmount = iso8583.NewAlphanumeric(basic.FormatTrxAmountString(msgInquiry.TransactionAmount))
		}
		isoFormat.ResponseCode = iso8583.NewAlphanumeric(msgInquiry.ResponseCode)

		log.Println("prepaid.IsoInquiry[Encode(message string)] : check the response code")
		if msgInquiry.ResponseCode == "0000" {
			isoFormat.AdditionalPrivateData =
				iso8583.NewLllvar([]byte(FormatInqRes(msgInquiry.AdditionalPrivateData.(*AdditionalPrivateData))))
			isoFormat.AdditionalPrivateData3 =
				iso8583.NewLllvar([]byte(FormatData3String(msgInquiry.AdditionalPrivateData3.(*AdditionalPrivateData3))))
		} else if msgInquiry.ResponseCode != "0000" {
			log.Println("prepaid.IsoInquiry[Encode(message string)] : response code not 0000")
			isoFormat.AdditionalPrivateData =
				iso8583.NewLllvar([]byte(FormatInqReq(msgInquiry.AdditionalPrivateData.(*AdditionalPrivateData))))
		}
	}

	msg := iso8583.NewMessageExtended(msgInquiry.Mti, iso8583.ASCII, false, true, isoFormat)

	packetIso, err := msg.Bytes()
	if err != nil {
		panic(err.Error())
	}

	return util.EncapsulateBytes(packetIso)

}

// Decode : decode from byte iso8583 to prepaid inquiry
func (isoInquiry *IsoInquiry) Decode(message []byte) (string, error) {

	log.Println("prepaid.IsoInquiry[Decode(message string)] : start to decode")
	resultFields, mti := basic.DecodeIsoMessage(message)

	log.Println("prepaid.IsoInquiry[Decode(message string)] : start to assign iso to message")
	msgInqResult := basic.AssignISOFormatToMessage(resultFields, mti)

	if mti == config.Get().Mti.Inquiry.Request {
		msgInqResult.AdditionalPrivateData = BuildInquiryReq(string(resultFields.AdditionalPrivateData.Value))
	} else if mti == config.Get().Mti.Inquiry.Response {
		msgInqResult.ResponseCode = resultFields.ResponseCode.Value
		if len(resultFields.TransactionAmount.Value) > 0 {
			msgInqResult.TransactionAmount = basic.ParseMessageToTrxAmt(resultFields.TransactionAmount.Value)
		}

		if resultFields.ResponseCode.Value != "0000" {
			msgInqResult.AdditionalPrivateData = BuildInquiryReq(
				string(resultFields.AdditionalPrivateData.Value))
		} else {
			msgInqResult.AdditionalPrivateData = BuildInquiryResponse(
				string(resultFields.AdditionalPrivateData.Value))

			log.Println("data private 3 : ", string(resultFields.AdditionalPrivateData3.Value))
			msgInqResult.AdditionalPrivateData3 = BuildData3(string(resultFields.AdditionalPrivateData3.Value))
		}
	}

	json, _ := json.Marshal(msgInqResult)

	return string(json), nil

}
