package isopostpaid

import (
	"encoding/json"

	log "gitlab.com/kasku/kasku-2pay/2pay-billerpayment/log"

	"github.com/Ayvan/iso8583"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/config"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/basic"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/util"
)

// IsoInquiry : iso inquiry
type IsoInquiry struct {
}

// Encode : to encode message for postpaid inquiry
func (isoInquiry *IsoInquiry) Encode(msgJSON string) []byte {
	log.Get().Println("postpaid.IsoInquiry[Encode(message string)] : start to encode ")

	log.Get().Println("postpaid.IsoInquiry[Encode(message string)] : initialize message to assign interface with isopostpaid message")
	message := &basic.Message{
		AdditionalPrivateData:  &AdditionalPrivateData{},
		AdditionalPrivateData2: &AdditionalPrivateData2{},
	}

	log.Get().Println("postpaid.IsoInquiry[Encode(message string)] : encode json format to iso")
	isoFormat, msgInquiry := basic.EncodeJSONFormatToISO(msgJSON, message)

	isoFormat.AdditionalPrivateData =
		iso8583.NewLllvar([]byte(FormatInquiryString(msgInquiry.AdditionalPrivateData.(*AdditionalPrivateData))))

	if msgInquiry.Mti == config.Get().Mti.Inquiry.Response {
		if len(msgInquiry.TransactionAmount.ValueAmount) > 0 {
			isoFormat.TransactionAmount = iso8583.NewAlphanumeric(basic.FormatTrxAmountString(msgInquiry.TransactionAmount))
		}
		isoFormat.ResponseCode = iso8583.NewAlphanumeric(msgInquiry.ResponseCode)

		if msgInquiry.ResponseCode == "0000" {
			isoFormat.AdditionalPrivateData =
				iso8583.NewLllvar([]byte(FormatDataString(msgInquiry.AdditionalPrivateData.(*AdditionalPrivateData))))
		}
	}

	msg := iso8583.NewMessageExtended(msgInquiry.Mti, iso8583.ASCII, false, true, isoFormat)

	packetIso, err := msg.Bytes()
	if err != nil {
		panic(err.Error())
	}

	return util.EncapsulateBytes(packetIso)

}

// Decode : decode from byte iso8583 to postpaid inquiry
func (isoInquiry *IsoInquiry) Decode(message []byte) (string, error) {

	log.Get().Println("postpaid.IsoInquiry[Decode(message string)] : start to decode")
	resultFields, mti := basic.DecodeIsoMessage(message)

	log.Get().Println("postpaid.IsoInquiry[Decode(message string)] : start to assign iso to message")
	msgInqResult := basic.AssignISOFormatToMessage(resultFields, mti)

	if mti == config.Get().Mti.Inquiry.Request {
		msgInqResult.AdditionalPrivateData = BuildInquiry(string(resultFields.AdditionalPrivateData.Value))
	} else if mti == config.Get().Mti.Inquiry.Response {
		msgInqResult.ResponseCode = resultFields.ResponseCode.Value
		if len(resultFields.TransactionAmount.Value) > 0 {
			msgInqResult.TransactionAmount = basic.ParseMessageToTrxAmt(resultFields.TransactionAmount.Value)
		}

		if resultFields.ResponseCode.Value != "0000" {
			msgInqResult.AdditionalPrivateData = BuildInquiry(string(resultFields.AdditionalPrivateData.Value))
		} else {
			msgInqResult.AdditionalPrivateData = ParsePostpaidAdditionalPrivateDataForInquiryResponse(
				string(resultFields.AdditionalPrivateData.Value))
		}
	}

	json, _ := json.Marshal(msgInqResult)

	return string(json), nil

}
