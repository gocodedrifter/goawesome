package processor

import (
	"encoding/json"
	"strings"

	log "gitlab.com/kasku/kasku-2pay/2pay-billerpayment/log"

	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/basic"
)

// JSONProcessor : json processor
type JSONProcessor struct {
}

// EncodeMessage : processing message from json format to iso bytes
func (jsonProcessor *JSONProcessor) EncodeMessage(message []byte) []byte {
	log.Get().Println("JsonProcessor[ProssesMessage(message []byte)] : start")
	msgType := &basic.MessageType{}
	if err := json.Unmarshal(message, msgType); err != nil {
		log.Get().Println("JSONProcessor[ProssesMessage(message []byte)] : unable to marshal")
	}

	buildIso := messaging.GetTypeMessage(strings.Join([]string{msgType.Mti, msgType.PrimaryAccountNumber}, ""))
	log.Get().Println("JSONProcessor[ProssesMessage(message []byte)] : get type message : ", strings.Join([]string{msgType.Mti, msgType.PrimaryAccountNumber}, ""))
	isobyte := messaging.EncodeMessage(buildIso, string(message))

	log.Get().Println("JSONProcessor[ProssesMessage(message []byte)] : result iso : ", string(isobyte))
	return isobyte
}

// DecodeMessage :
func (jsonProcessor *JSONProcessor) DecodeMessage(message []byte) string {
	return ""
}
