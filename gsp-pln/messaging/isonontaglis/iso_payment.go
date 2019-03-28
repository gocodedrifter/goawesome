package isonontaglis

import (
	"encoding/json"
	"log"

	"github.com/Ayvan/iso8583"

	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/basic"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/util"
)

// IsoPayment : non tagihan listrik payment
type IsoPayment struct {
}

// Encode : to encode message for nontaglis payment
func (isoPayment *IsoPayment) Encode(message string) []byte {

	log.Println("nontaglis.IsoInquiry[Encode(message string)] : start to encode")

	msgPayment := &Message{}
	if err := json.Unmarshal([]byte(message), msgPayment); err != nil {
		log.Println("nontaglis.IsoInquiry[Encode(message string)] : unable to marshal")
	}

	msg := iso8583.NewMessageExtended(msgPayment.Mti, iso8583.ASCII, false, true,
		&basic.Iso8583Format{
			PrimaryAccountNumber:     iso8583.NewLlvar([]byte(msgPayment.PrimaryAccountNumber)),
			TransactionAmount:        iso8583.NewAlphanumeric(basic.FormatTrxAmountString(msgPayment.TransactionAmount)),
			Stan:                     iso8583.NewAlphanumeric(util.GetIsoStanFormat(msgPayment.Stan)),
			DateTimeLocalTransaction: iso8583.NewAlphanumeric(msgPayment.DateTimeLocalTransaction),
			MerchantCategoryCode:     iso8583.NewAlphanumeric(msgPayment.MerchantCategoryCode),
			BankCode:                 iso8583.NewLlvar([]byte(util.GetIsoBankCodeFormat(msgPayment.BankCode))),
			PartnerCentralID:         iso8583.NewLlvar([]byte(msgPayment.PartnerCentralID)),
			TerminalID:               iso8583.NewAlphanumeric(util.GetIsoTerminalIDFormat(msgPayment.TerminalID)),
			AdditionalPrivateData:    iso8583.NewLllvar([]byte(FormatDataString(msgPayment.AdditionalPrivateData))),
			AdditionalPrivateData3:   iso8583.NewLllvar([]byte(FormatData3String(msgPayment.AdditionalPrivateData3))),
		})

	packetIso, err := msg.Bytes()
	if err != nil {
		panic(err.Error())
	}

	return util.EncapsulateBytes(packetIso)
}

// Decode : decode from byte iso8583 to nontaglis payment
func (isoPayment *IsoPayment) Decode(message []byte) (string, error) {

	log.Println("nontaglis.IsoInquiry[Decode(message string)] : start to decode")
	resultFields := basic.DecodeIsoMessage(message)

	msgInqResult := &Message{
		PrimaryAccountNumber:     string(resultFields.PrimaryAccountNumber.Value),
		TransactionAmount:        basic.ParseMessageToTrxAmt(resultFields.TransactionAmount.Value),
		Stan:                     resultFields.Stan.Value,
		SettlementDate:           resultFields.SettlementDate.Value,
		DateTimeLocalTransaction: resultFields.DateTimeLocalTransaction.Value,
		MerchantCategoryCode:     resultFields.MerchantCategoryCode.Value,
		BankCode:                 string(resultFields.BankCode.Value),
		PartnerCentralID:         string(resultFields.PartnerCentralID.Value),
		ResponseCode:             resultFields.ResponseCode.Value,
		TerminalID:               resultFields.TerminalID.Value,
		AdditionalPrivateData:    BuildResponse(string(resultFields.AdditionalPrivateData.Value)),
		AdditionalPrivateData2:   BuildData2Response(string(resultFields.AdditionalPrivateData2.Value)),
		AdditionalPrivateData3:   BuildData3Respose(string(resultFields.AdditionalPrivateData3.Value)),
	}

	json, _ := json.Marshal(msgInqResult)

	return string(json), nil
}
