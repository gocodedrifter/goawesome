package isonontaglis

import (
	"encoding/json"
	"log"

	"github.com/Ayvan/iso8583"

	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/basic"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/util"
)

// IsoInquiry : non tagihan listrik inquiry
type IsoInquiry struct {
}

// Encode : to encode message for nontaglis inquiry
func (isoInquiry *IsoInquiry) Encode(message string) []byte {

	log.Println("nontaglis.IsoInquiry[Encode(message string)] : start to encode")

	msgInquiry := &Message{}
	if err := json.Unmarshal([]byte(message), msgInquiry); err != nil {
		log.Println("nontaglis.IsoInquiry[Encode(message string)] : unable to marshal")
	}

	msg := iso8583.NewMessageExtended(msgInquiry.Mti, iso8583.ASCII, false, true,
		&basic.Iso8583Format{
			PrimaryAccountNumber:     iso8583.NewLlvar([]byte(msgInquiry.PrimaryAccountNumber)),
			Stan:                     iso8583.NewAlphanumeric(util.GetIsoStanFormat(msgInquiry.Stan)),
			DateTimeLocalTransaction: iso8583.NewAlphanumeric(msgInquiry.DateTimeLocalTransaction),
			MerchantCategoryCode:     iso8583.NewAlphanumeric(msgInquiry.MerchantCategoryCode),
			BankCode:                 iso8583.NewLlvar([]byte(util.GetIsoBankCodeFormat(msgInquiry.BankCode))),
			PartnerCentralID:         iso8583.NewLlvar([]byte(msgInquiry.PartnerCentralID)),
			TerminalID:               iso8583.NewAlphanumeric(util.GetIsoTerminalIDFormat(msgInquiry.TerminalID)),
			AdditionalPrivateData:    iso8583.NewLllvar([]byte(FormatInquiryString(msgInquiry.AdditionalPrivateData))),
		})

	packetIso, err := msg.Bytes()
	if err != nil {
		panic(err.Error())
	}

	return util.EncapsulateBytes(packetIso)
}

// Decode : decode from byte iso8583 to nontaglis inquiry
func (isoInquiry *IsoInquiry) Decode(message []byte) (string, error) {

	log.Println("nontaglis.IsoInquiry[Decode(message string)] : start to decode")
	resultFields := basic.DecodeIsoMessage(message)

	msgInqResult := &Message{
		PrimaryAccountNumber:     string(resultFields.PrimaryAccountNumber.Value),
		TransactionAmount:        basic.ParseMessageToTrxAmt(resultFields.TransactionAmount.Value),
		Stan:                     resultFields.Stan.Value,
		DateTimeLocalTransaction: resultFields.DateTimeLocalTransaction.Value,
		MerchantCategoryCode:     resultFields.MerchantCategoryCode.Value,
		BankCode:                 string(resultFields.BankCode.Value),
		PartnerCentralID:         string(resultFields.PartnerCentralID.Value),
		ResponseCode:             resultFields.ResponseCode.Value,
		TerminalID:               resultFields.TerminalID.Value,
	}

	if resultFields.ResponseCode.Value != "0000" {
		msgInqResult.AdditionalPrivateData = BuildResponseUnexpected(string(resultFields.AdditionalPrivateData.Value))
	} else {
		msgInqResult.AdditionalPrivateData = BuildResponse(string(resultFields.AdditionalPrivateData.Value))
		msgInqResult.AdditionalPrivateData3 = BuildData3Respose(string(resultFields.AdditionalPrivateData3.Value))
	}

	json, _ := json.Marshal(msgInqResult)

	return string(json), nil
}
