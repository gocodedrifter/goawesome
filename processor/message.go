package processor

import log "gitlab.com/kasku/kasku-2pay/2pay-billerpayment/log"

// Message : message
type Message struct {
	buildProcess BuildProcess
}

// SetBuilder : set type builder
func (message *Message) SetBuilder(buildProcess BuildProcess) {
	log.Get().Println("message[SetBuilder] : ", buildProcess)
	message.buildProcess = buildProcess
}

// Process : processing the data
func (message *Message) Process(data []byte) []byte {
	return message.buildProcess.EncodeMessage(data)
}

// DecodeMessage : encode the message
func (message *Message) DecodeMessage(data []byte) string {
	return message.buildProcess.DecodeMessage(data)
}
