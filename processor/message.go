package processor

// Message : message
type Message struct {
	buildProcess BuildProcess
}

// SetBuilder : set type builder
func (message *Message) SetBuilder(buildProcess BuildProcess) {
	message.buildProcess = buildProcess
}

// Process : processing the data
func (message *Message) Process(data []byte) []byte {
	return message.buildProcess.ProssesMessage(data)
}

// EncodeMessage : encode the message
func (message *Message) EncodeMessage(data []byte) string {
	return message.buildProcess.EncodeMessage(data)
}
