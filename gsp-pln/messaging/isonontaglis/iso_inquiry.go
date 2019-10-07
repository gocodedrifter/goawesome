package isonontaglis

import (
	"encoding/json"

	log "gitlab.com/kasku/kasku-2pay/2pay-billerpayment/log"

	"github.com/Ayvan/iso8583"

	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/config"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/basic"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/util"
)

// IsoInquiry : non tagihan listrik inquiry
type IsoInquiry struct {
}

// Encode : to encode message for nontaglis inquiry
func (isoInquiry *IsoInquiry) Encode(msgJSON string) []byte {

	log.Get().Println("nontaglis.IsoInquiry[Encode(message string)] : start to encode ")

	log.Get().Println("nontaglis.IsoInquiry[Encode(message string)] : initialize message to assign interface with isopostpaid message")
	message := &basic.Message{
		AdditionalPrivateData:  &AdditionalPrivateData{},
		AdditionalPrivateData2: &AdditionalPrivateData2{},
		AdditionalPrivateData3: &AdditionalPrivateData3{},
	}

	log.Get().Println("nontaglis.IsoInquiry[Encode(message string)] : encode json format to iso")
	isoFormat, msgInquiry := basic.EncodeJSONFormatToISO(msgJSON, message)

	if msgInquiry.Mti == config.Get().Mti.Inquiry.Request {
		isoFormat.AdditionalPrivateData = iso8583.NewLllvar([]byte(FormatInquiryString(msgInquiry.AdditionalPrivateData.(*AdditionalPrivateData))))
	} else if msgInquiry.Mti == config.Get().Mti.Inquiry.Response {
		if len(msgInquiry.TransactionAmount.ValueAmount) > 0 {
			isoFormat.TransactionAmount = iso8583.NewAlphanumeric(basic.FormatTrxAmountString(msgInquiry.TransactionAmount))
		}
		isoFormat.ResponseCode = iso8583.NewAlphanumeric(msgInquiry.ResponseCode)
		isoFormat.AdditionalPrivateData = iso8583.NewLllvar([]byte(FormatDataString(msgInquiry.AdditionalPrivateData.(*AdditionalPrivateData))))
		if msgInquiry.ResponseCode == "0000" {
			isoFormat.AdditionalPrivateData2 = iso8583.NewLllvar([]byte(FormatData2String(msgInquiry.AdditionalPrivateData2.(*AdditionalPrivateData2))))
			isoFormat.AdditionalPrivateData3 = iso8583.NewLllvar([]byte(FormatData3String(msgInquiry.AdditionalPrivateData3.(*AdditionalPrivateData3))))
		}
	}

	msg := iso8583.NewMessageExtended(msgInquiry.Mti, iso8583.ASCII, false, true, isoFormat)

	packetIso, err := msg.Bytes()
	if err != nil {
		panic(err.Error())
	}

	return util.EncapsulateBytes(packetIso)
}

// Decode : decode from byte iso8583 to nontaglis inquiry
func (isoInquiry *IsoInquiry) Decode(message []byte) (string, error) {

	log.Get().Println("nontaglis.IsoInquiry[Decode(message string)] : start to decode")
	resultFields, mti := basic.DecodeIsoMessage(message)

	log.Get().Println("nontaglis.IsoInquiry[Decode(message string)] : start to assign iso to message")
	msgInqResult := basic.AssignISOFormatToMessage(resultFields, mti)

	if mti == config.Get().Mti.Inquiry.Request {
		msgInqResult.AdditionalPrivateData = BuildInquiry(string(resultFields.AdditionalPrivateData.Value))
	} else if mti == config.Get().Mti.Inquiry.Response {
		msgInqResult.ResponseCode = resultFields.ResponseCode.Value
		msgInqResult.TransactionAmount = basic.ParseMessageToTrxAmt(resultFields.TransactionAmount.Value)
		if resultFields.ResponseCode.Value != "0000" {
			msgInqResult.AdditionalPrivateData = BuildInquiry(string(resultFields.AdditionalPrivateData.Value))
		} else {
			msgInqResult.AdditionalPrivateData = BuildResponse(string(resultFields.AdditionalPrivateData.Value))
			msgInqResult.AdditionalPrivateData3 = BuildData3Respose(string(resultFields.AdditionalPrivateData3.Value))
		}
	}

	json, _ := json.Marshal(msgInqResult)

	return string(json), nil
}
