package processor

import (
	"bytes"
	"strings"

	log "gitlab.com/kasku/kasku-2pay/2pay-billerpayment/log"

	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging"
)

// IsoProcessor : iso processor
type IsoProcessor struct {
}

// EncodeMessage : process message
func (isoProcessor *IsoProcessor) EncodeMessage(message []byte) []byte {
	return nil
}

// DecodeMessage : decode message from isobytes to json string
func (isoProcessor *IsoProcessor) DecodeMessage(message []byte) string {
	var mti bytes.Buffer
	log.Get().Println("isoProcessor[DecodeMessage] received message : ", string(message))
	log.Get().Println("check message : ", string(message[strings.Index(string(message), "2"):strings.Index(string(message), "2")+4]))
	log.Get().Println("check primary account number : ", string(message[strings.Index(string(message), "2")+22:strings.Index(string(message), "2")+27]))
	if mti.WriteString(string(message[strings.Index(string(message), "2") : strings.Index(string(message), "2")+4])); !strings.HasPrefix(mti.String(), "28") {
		mti.WriteString(string(message[strings.Index(string(message), "2")+22 : strings.Index(string(message), "2")+27]))
	}
	buildIso := messaging.GetTypeMessage(mti.String())
	log.Get().Println("build iso : ", buildIso)
	jsonResult, _ := messaging.DecodeMessage(buildIso, message)
	log.Get().Println("isoProcessor[DecodeMessage] result : ", jsonResult)
	return jsonResult
}
