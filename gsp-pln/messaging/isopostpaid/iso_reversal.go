package isopostpaid

import (
	"encoding/json"
	"log"

	"github.com/Ayvan/iso8583"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/config"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/basic"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/util"
)

// IsoReversal : iso reversal for postpaid
type IsoReversal struct {
}

// Encode : to encode message for postpaid reversal
func (isoReversal *IsoReversal) Encode(msgJSON string) []byte {

	log.Println("postpaid.IsoReversal[Encode(message string)] : start to encode")

	message := &basic.Message{
		AdditionalPrivateData:  &AdditionalPrivateData{},
		AdditionalPrivateData2: &AdditionalPrivateData2{},
	}

	log.Println("postpaid.IsoInquiry[Encode(message string)] : encode json format to iso")
	isoFormat, msgReversal := basic.EncodeJSONFormatToISO(msgJSON, message)

	isoFormat.AdditionalPrivateData =
		iso8583.NewLllvar([]byte(FormatDataString(msgReversal.AdditionalPrivateData.(*AdditionalPrivateData))))
	isoFormat.OriginalData = iso8583.NewLlvar([]byte(basic.FormatReversalString(msgReversal.OriginalData)))
	isoFormat.TransactionAmount = iso8583.NewAlphanumeric(basic.FormatTrxAmountString(msgReversal.TransactionAmount))
	if msgReversal.Mti == config.Get().Mti.Reversal.Response || msgReversal.Mti == config.Get().Mti.Reversal.Repeat.Response {
		isoFormat.ResponseCode = iso8583.NewAlphanumeric(msgReversal.ResponseCode)
	}

	msg := iso8583.NewMessageExtended(msgReversal.Mti, iso8583.ASCII, false, true, isoFormat)

	packetIso, err := msg.Bytes()
	if err != nil {
		panic(err.Error())
	}

	return util.EncapsulateBytes(packetIso)
}

// Decode : decode from byte iso8583 to postpaid reversal
func (isoReversal *IsoReversal) Decode(message []byte) (string, error) {

	log.Println("postpaid.IsoReversal[Decode(message string)] : start to decode")
	resultFields, mti := basic.DecodeIsoMessage(message)

	log.Println("postpaid.IsoInquiry[Decode(message string)] : start to assign iso to message")
	msgReversal := basic.AssignISOFormatToMessage(resultFields, mti)

	msgReversal.AdditionalPrivateData = BuildDataResponse(string(resultFields.AdditionalPrivateData.Value))
	msgReversal.OriginalData = basic.BuildOriginalDataResponse(string(resultFields.OriginalData.Value))
	msgReversal.TransactionAmount = basic.ParseMessageToTrxAmt(resultFields.TransactionAmount.Value)

	if len(resultFields.ResponseCode.Value) > 0 {
		msgReversal.ResponseCode = resultFields.ResponseCode.Value
	}

	json, _ := json.Marshal(msgReversal)

	return string(json), nil
}
