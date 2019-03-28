package isonetman

import (
	"encoding/json"
	"log"

	"github.com/Ayvan/iso8583"
	"gitlab.com/2pay/biller-payment/utils"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/basic"
)

// Netman : network management
type Netman struct {
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

	netmanMsg := Netman{}
	if err := json.Unmarshal([]byte(message), &netmanMsg); err != nil {
		log.Println("netman[Encode(message string)] : unable to marshal")
	}

	msg := iso8583.NewMessageExtended(utils.MtiNetmanRequest, iso8583.ASCII, false, true,
		&basic.Iso8583Format{
			DateTimeLocalTransaction: iso8583.NewAlphanumeric(netmanMsg.TransactionTime),
			PartnerCentralID:         iso8583.NewLlvar([]byte(netmanMsg.PartnerCentralID)),
			ActionCode:               iso8583.NewAlphanumeric(netmanMsg.ActionCode),
			TerminalID:               iso8583.NewAlphanumeric(netmanMsg.TerminalID),
			AdditionalPrivateData:    iso8583.NewLllvar([]byte(FormatString(netmanMsg.AdditionalPrivateData))),
		})

	packetIso, err := msg.Bytes()
	if err != nil {
		panic(err.Error())
	}

	return utils.EncapsulateBytes(packetIso)

}

// Decode : decode from byte iso8583 to networkmanagement
func (netman *Netman) Decode(message []byte) (string, error) {

	iso := iso8583.NewMessageExtended("", iso8583.ASCII, false, true,
		&basic.Iso8583Format{
			DateTimeLocalTransaction: iso8583.NewAlphanumeric(""),
			PartnerCentralID:         iso8583.NewLlvar([]byte("")),
			ActionCode:               iso8583.NewAlphanumeric(""),
			ResponseCode:             iso8583.NewAlphanumeric(""),
			TerminalID:               iso8583.NewAlphanumeric(""),
			AdditionalPrivateData:    iso8583.NewLllvar([]byte("")),
		})

	if err := iso.Load(message[2:]); err != nil {
		panic(err.Error())
	}

	resultFields := iso.Data.(*basic.Iso8583Format)

	netmanResult := &Netman{
		TransactionTime:       resultFields.DateTimeLocalTransaction.Value,
		PartnerCentralID:      string(resultFields.PartnerCentralID.Value),
		ActionCode:            resultFields.ActionCode.Value,
		TerminalID:            resultFields.TerminalID.Value,
		ResponseCode:          resultFields.ResponseCode.Value,
		AdditionalPrivateData: BuildAdditionalPrivateData(string(resultFields.AdditionalPrivateData.Value)),
	}

	json, _ := json.Marshal(netmanResult)

	return string(json), nil
}
