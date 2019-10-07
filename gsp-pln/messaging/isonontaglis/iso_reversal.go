package isonontaglis

import (
	"encoding/json"

	log "gitlab.com/kasku/kasku-2pay/2pay-billerpayment/log"

	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/config"

	"github.com/Ayvan/iso8583"

	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/basic"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/util"
)

// IsoReversal : reversal non tagihan listrik
type IsoReversal struct {
}

// Encode : to encode message for nontaglis reversal
func (isoReversal *IsoReversal) Encode(msgJSON string) []byte {

	log.Get().Println("postpaid.IsoReversal[Encode(message string)] : start to encode")

	message := &basic.Message{
		AdditionalPrivateData:  &AdditionalPrivateData{},
		AdditionalPrivateData2: &AdditionalPrivateData2{},
		AdditionalPrivateData3: &AdditionalPrivateData3{},
	}

	log.Get().Println("postpaid.IsoInquiry[Encode(message string)] : encode json format to iso")
	isoFormat, msgReversal := basic.EncodeJSONFormatToISO(msgJSON, message)

	isoFormat.TransactionAmount = iso8583.NewAlphanumeric(basic.FormatTrxAmountString(msgReversal.TransactionAmount))
	isoFormat.AdditionalPrivateData = iso8583.NewLllvar([]byte(FormatDataString(msgReversal.AdditionalPrivateData.(*AdditionalPrivateData))))
	isoFormat.OriginalData = iso8583.NewLlvar([]byte(basic.FormatReversalString(msgReversal.OriginalData)))
	isoFormat.AdditionalPrivateData3 = iso8583.NewLllvar([]byte(FormatData3String(msgReversal.AdditionalPrivateData3.(*AdditionalPrivateData3))))

	if msgReversal.Mti == config.Get().Mti.Reversal.Response || msgReversal.Mti == config.Get().Mti.Reversal.Repeat.Response {
		isoFormat.ResponseCode = iso8583.NewAlphanumeric(msgReversal.ResponseCode)
	}

	if len(msgReversal.AdditionalPrivateData2.(*AdditionalPrivateData2).PowerConsumingCategory) > 0 {
		isoFormat.AdditionalPrivateData2 = iso8583.NewLllvar([]byte(FormatData2String(msgReversal.AdditionalPrivateData2.(*AdditionalPrivateData2))))
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

	log.Get().Println("nontaglis.IsoReversal[Decode(message string)] : start to decode")
	resultFields, mti := basic.DecodeIsoMessage(message)

	log.Get().Println("nontaglis.IsoReversal[Decode(message string)] : start to assign iso to message")
	msgReversal := basic.AssignISOFormatToMessage(resultFields, mti)

	msgReversal.TransactionAmount = basic.ParseMessageToTrxAmt(resultFields.TransactionAmount.Value)
	msgReversal.AdditionalPrivateData = BuildResponse(string(resultFields.AdditionalPrivateData.Value))
	msgReversal.OriginalData = basic.BuildOriginalDataResponse(string(resultFields.OriginalData.Value))
	// AdditionalPrivateData2:   BuildData2Response(string(resultFields.AdditionalPrivateData2.Value)),
	msgReversal.AdditionalPrivateData3 = BuildData3Respose(string(resultFields.AdditionalPrivateData3.Value))

	if mti == config.Get().Mti.Reversal.Response || mti == config.Get().Mti.Reversal.Repeat.Response {
		msgReversal.ResponseCode = resultFields.ResponseCode.Value
	}

	json, _ := json.Marshal(msgReversal)

	return string(json), nil
}
