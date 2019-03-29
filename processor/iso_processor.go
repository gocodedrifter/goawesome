package processor

import (
	"bytes"

	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/config"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging"
)

// IsoProcessor : iso processor
type IsoProcessor struct {
}

// ProssesMessage : process message
func (isoProcessor *IsoProcessor) ProssesMessage(message []byte) []byte {
	return nil
}

// EncodeMessage : encode the message
func (isoProcessor *IsoProcessor) EncodeMessage(message []byte) string {
	var mti bytes.Buffer
	if mti.WriteString(string(message[2:6])); mti.String() != config.Get().Mti.Netman.Response {
		mti.WriteString(string(message[24:29]))
	}
	buildIso := messaging.GetTypeMessage(mti.String())
	jsonResult, _ := buildIso.Decode(message)
	return jsonResult
}
