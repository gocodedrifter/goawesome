package isonontaglis

import (
	"encoding/json"
	"log"

	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/config"

	"github.com/Ayvan/iso8583"

	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/basic"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/util"
)

// IsoReversal : reversal non tagihan listrik
type IsoReversal struct {
}

// Encode : to encode message for nontaglis reversal
func (isoReversal *IsoReversal) Encode(message string) []byte {

	log.Println("nontaglis.IsoReversal[Encode(message string)] : start to encode")

	msgReversal := &Message{}
	if err := json.Unmarshal([]byte(message), msgReversal); err != nil {
		log.Println("nontaglis.IsoReversal[Encode(message string)] : unable to marshal")
	}

	isoFormat := &basic.Iso8583Format{
		PrimaryAccountNumber:     iso8583.NewLlvar([]byte(msgReversal.PrimaryAccountNumber)),
		TransactionAmount:        iso8583.NewAlphanumeric(basic.FormatTrxAmountString(msgReversal.TransactionAmount)),
		Stan:                     iso8583.NewAlphanumeric(util.GetIsoStanFormat(msgReversal.Stan)),
		DateTimeLocalTransaction: iso8583.NewAlphanumeric(msgReversal.DateTimeLocalTransaction),
		MerchantCategoryCode:     iso8583.NewAlphanumeric(msgReversal.MerchantCategoryCode),
		BankCode:                 iso8583.NewLlvar([]byte(util.GetIsoBankCodeFormat(msgReversal.BankCode))),
		PartnerCentralID:         iso8583.NewLlvar([]byte(msgReversal.PartnerCentralID)),
		TerminalID:               iso8583.NewAlphanumeric(util.GetIsoTerminalIDFormat(msgReversal.TerminalID)),
		AdditionalPrivateData:    iso8583.NewLllvar([]byte(FormatDataString(msgReversal.AdditionalPrivateData))),
		OriginalData:             iso8583.NewLlvar([]byte(FormatReversalString(msgReversal.OriginalData))),
		// AdditionalPrivateData2:   iso8583.NewLllvar([]byte(FormatData2String(msgReversal.AdditionalPrivateData2))),
		AdditionalPrivateData3: iso8583.NewLllvar([]byte(FormatData3String(msgReversal.AdditionalPrivateData3))),
	}

	if msgReversal.Mti == config.Get().Mti.Reversal.Response {
		isoFormat.ResponseCode = iso8583.NewAlphanumeric(msgReversal.ResponseCode)
	}

	if len(msgReversal.AdditionalPrivateData2.PowerConsumingCategory) > 0 {
		isoFormat.AdditionalPrivateData2 = iso8583.NewLllvar([]byte(FormatData2String(msgReversal.AdditionalPrivateData2)))
	}

	msg := iso8583.NewMessageExtended(msgReversal.Mti, iso8583.ASCII, false, true, isoFormat)

	packetIso, err := msg.Bytes()
	if err != nil {
		panic(err.Error())
	}

	return util.EncapsulateBytes(packetIso)
}

// Decode : decode from byte iso8583 to nontaglis reversal
func (isoReversal *IsoReversal) Decode(message []byte) (string, error) {

	log.Println("nontaglis.IsoReversal[Decode(message string)] : start to decode")
	resultFields, mti := basic.DecodeIsoMessage(message)

	msgReversal := &Message{
		Mti: mti,
		// Payload:                  string(message[2:]),
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
		OriginalData:             BuildOriginalDataResponse(string(resultFields.OriginalData.Value)),
		// AdditionalPrivateData2:   BuildData2Response(string(resultFields.AdditionalPrivateData2.Value)),
		AdditionalPrivateData3: BuildData3Respose(string(resultFields.AdditionalPrivateData3.Value)),
	}

	if mti == config.Get().Mti.Reversal.Response {
		msgReversal.ResponseCode = resultFields.ResponseCode.Value
	}

	json, _ := json.Marshal(msgReversal)

	return string(json), nil
}
