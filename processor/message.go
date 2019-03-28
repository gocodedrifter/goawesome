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
	return message.Process(data)
}
