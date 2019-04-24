package isonontaglis

import (
	"encoding/json"
	"log"

	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/config"

	"github.com/Ayvan/iso8583"

	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/basic"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/util"
)

// IsoPayment : non tagihan listrik payment
type IsoPayment struct {
}

// Encode : to encode message for nontaglis payment
func (isoPayment *IsoPayment) Encode(message string) []byte {

	log.Println("nontaglis.IsoPayment[Encode(message string)] : start to encode")

	msgPayment := &Message{}
	if err := json.Unmarshal([]byte(message), msgPayment); err != nil {
		log.Println("nontaglis.IsoPayment[Encode(message string)] : unable to marshal")
	}

	isoFormat := &basic.Iso8583Format{
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
	}

	if msgPayment.Mti == config.Get().Mti.Payment.Response {
		isoFormat.ResponseCode = iso8583.NewAlphanumeric(msgPayment.ResponseCode)
		if msgPayment.ResponseCode == "0000" {
			isoFormat.AdditionalPrivateData2 = iso8583.NewLllvar([]byte(FormatData2String(msgPayment.AdditionalPrivateData2)))
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

	log.Println("nontaglis.IsoPayment[Decode(message string)] : start to decode")
	resultFields, mti := basic.DecodeIsoMessage(message)

	msgPayResult := &Message{
		Mti:                      mti,
		PrimaryAccountNumber:     string(resultFields.PrimaryAccountNumber.Value),
		TransactionAmount:        basic.ParseMessageToTrxAmt(resultFields.TransactionAmount.Value),
		Stan:                     resultFields.Stan.Value,
		SettlementDate:           resultFields.SettlementDate.Value,
		DateTimeLocalTransaction: resultFields.DateTimeLocalTransaction.Value,
		MerchantCategoryCode:     resultFields.MerchantCategoryCode.Value,
		BankCode:                 string(resultFields.BankCode.Value),
		PartnerCentralID:         string(resultFields.PartnerCentralID.Value),
		TerminalID:               resultFields.TerminalID.Value,
		AdditionalPrivateData:    BuildResponse(string(resultFields.AdditionalPrivateData.Value)),
		AdditionalPrivateData3:   BuildData3Respose(string(resultFields.AdditionalPrivateData3.Value)),
	}

	if mti == config.Get().Mti.Payment.Response {
		msgPayResult.ResponseCode = resultFields.ResponseCode.Value
		if resultFields.ResponseCode.Value == "0000" {
			msgPayResult.AdditionalPrivateData2 = BuildData2Response(string(resultFields.AdditionalPrivateData2.Value))
		}
	}

	log.Println("nontaglis.IsoPayment[Decode(message string)] json decode : ", &msgPayResult)

	json, _ := json.Marshal(msgPayResult)

	return string(json), nil
}
