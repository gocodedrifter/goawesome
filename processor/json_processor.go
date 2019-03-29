package processor

import (
	"encoding/json"
	"log"
	"strings"

	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/basic"
)

// JSONProcessor : json processor
type JSONProcessor struct {
}

// ProssesMessage : processing message for json
func (jsonProcessor *JSONProcessor) ProssesMessage(message []byte) []byte {
	msgType := &basic.MessageType{}
	if err := json.Unmarshal(message, msgType); err != nil {
		log.Println("JSONProcessor[ProssesMessage(message []byte)] : unable to marshal")
	}

	buildIso := messaging.GetTypeMessage(strings.Join([]string{msgType.Mti, msgType.PrimaryAccountNumber}, ""))
	log.Println("SONProcessor[ProssesMessage(message []byte)] : get type message : ", strings.Join([]string{msgType.Mti, msgType.PrimaryAccountNumber}, ""))
	isobyte := messaging.EncodeMessage(buildIso, string(message))

	return isobyte
}

// EncodeMessage : encode message iso bytes to json
func (jsonProcessor *JSONProcessor) EncodeMessage(message []byte) string {
	return ""
}
