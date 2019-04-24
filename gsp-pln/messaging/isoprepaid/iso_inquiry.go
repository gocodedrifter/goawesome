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
func (isoInquiry *IsoInquiry) Encode(message string) []byte {
	log.Println("prepaid.IsoInquiry[Encode(message string)] : start to encode ")

	msgInquiry := &Message{}
	if err := json.Unmarshal([]byte(message), msgInquiry); err != nil {
		log.Println("prepaid.IsoInquiry[Encode(message string)] : unable to marshal")
	}

	log.Println("prepaid.IsoInquiry[Encode(message string)] : mti to encode ", msgInquiry.Mti)
	isoFormat := &basic.Iso8583Format{
		PrimaryAccountNumber:     iso8583.NewLlvar([]byte(msgInquiry.PrimaryAccountNumber)),
		Stan:                     iso8583.NewAlphanumeric(util.GetIsoStanFormat(msgInquiry.Stan)),
		DateTimeLocalTransaction: iso8583.NewAlphanumeric(msgInquiry.DateTimeLocalTransaction),
		MerchantCategoryCode:     iso8583.NewAlphanumeric(msgInquiry.MerchantCategoryCode),
		BankCode:                 iso8583.NewLlvar([]byte(util.GetIsoBankCodeFormat(msgInquiry.BankCode))),
		PartnerCentralID:         iso8583.NewLlvar([]byte(msgInquiry.PartnerCentralID)),
		TerminalID:               iso8583.NewAlphanumeric(util.GetIsoTerminalIDFormat(msgInquiry.TerminalID)),
	}

	if msgInquiry.Mti == config.Get().Mti.Inquiry.Request {
		isoFormat.AdditionalPrivateData = iso8583.NewLllvar([]byte(FormatInquiryString(msgInquiry.AdditionalPrivateData)))
	} else if msgInquiry.Mti == config.Get().Mti.Inquiry.Response {
		if len(msgInquiry.TransactionAmount.ValueAmount) > 0 {
			isoFormat.TransactionAmount = iso8583.NewAlphanumeric(basic.FormatTrxAmountString(msgInquiry.TransactionAmount))
		}
		isoFormat.ResponseCode = iso8583.NewAlphanumeric(msgInquiry.ResponseCode)

		if msgInquiry.ResponseCode == "0000" {
			isoFormat.AdditionalPrivateData = iso8583.NewLllvar([]byte(FormatDataString(msgInquiry.AdditionalPrivateData)))
			isoFormat.AdditionalPrivateData3 = iso8583.NewLllvar([]byte(FormatData3String(msgInquiry.AdditionalPrivateData3)))
		}
	}

	msg := iso8583.NewMessageExtended(msgInquiry.Mti, iso8583.ASCII, false, true, isoFormat)

	packetIso, err := msg.Bytes()
	if err != nil {
		panic(err.Error())
	}

	return util.EncapsulateBytes(packetIso)

}
