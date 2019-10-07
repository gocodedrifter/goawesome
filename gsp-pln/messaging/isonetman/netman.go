package isonetman

import (
	"encoding/json"

	log "gitlab.com/kasku/kasku-2pay/2pay-billerpayment/log"

	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/config"

	"github.com/Ayvan/iso8583"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/basic"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/util"
)

// Netman : network management
type Netman struct {
	Mti                   string                `json:"mti,omitempty"`
	Payload               string                `json:"payload,omitempty"`
	TransactionTime       string                `json:"transactionTime,omitempty"`
	PartnerCentralID      string                `json:"partnerCentralId,omitempty"`
	ActionCode            string                `json:"actionCode,omitempty"`
	ResponseCode          string                `json:"responseCode,omitempty"`
	TerminalID            string                `json:"terminalId,omitempty"`
	AdditionalPrivateData AdditionalPrivateData `json:"additionalPrivateData,omitempty"`
}

// Encode : encode from network management to iso8583
func (netman *Netman) Encode(message string) []byte {

	log.Get().Println("netman[Encode(message string)] : start to encode")
	netmanMsg := Netman{}
	if err := json.Unmarshal([]byte(message), &netmanMsg); err != nil {
		log.Get().Println("netman[Encode(message string)] : unable to marshal")
	}

	isoFormat := &basic.Iso8583Format{
		DateTimeLocalTransaction: iso8583.NewAlphanumeric(netmanMsg.TransactionTime),
		PartnerCentralID:         iso8583.NewLlvar([]byte(netmanMsg.PartnerCentralID)),
		ActionCode:               iso8583.NewAlphanumeric(netmanMsg.ActionCode),
		TerminalID:               iso8583.NewAlphanumeric(util.GetIsoTerminalIDFormat(netmanMsg.TerminalID)),
		AdditionalPrivateData:    iso8583.NewLllvar([]byte(FormatString(netmanMsg.AdditionalPrivateData))),
	}

	if netmanMsg.Mti == config.Get().Mti.Netman.Response {
		isoFormat.ResponseCode = iso8583.NewAlphanumeric(netmanMsg.ResponseCode)
	}

	msg := iso8583.NewMessageExtended(netmanMsg.Mti, iso8583.ASCII, false, true, isoFormat)

	packetIso, err := msg.Bytes()
	if err != nil {
		panic(err.Error())
	}

	log.Get().Println("netman[Encode(message string)] : end to encode")
	return util.EncapsulateBytes(packetIso)

}

// Decode : decode from byte iso8583 to networkmanagement
func (netman *Netman) Decode(message []byte) (string, error) {

	log.Get().Println("netman[Decode(message string)] : start to decode")
	resultFields, mti := basic.DecodeIsoMessage(message)

	netmanResult := &Netman{
		Mti:                   mti,
		TransactionTime:       resultFields.DateTimeLocalTransaction.Value,
		PartnerCentralID:      string(resultFields.PartnerCentralID.Value),
		ActionCode:            resultFields.ActionCode.Value,
		TerminalID:            resultFields.TerminalID.Value,
		ResponseCode:          resultFields.ResponseCode.Value,
		AdditionalPrivateData: BuildAdditionalPrivateData(string(resultFields.AdditionalPrivateData.Value)),
	}

	result, _ := json.Marshal(netmanResult)

	log.Get().Println("netman[Decode(message string)] : end to decode")
	return string(result), nil
}
