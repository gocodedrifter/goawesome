package processor

// BuildProcess : build process message
type BuildProcess interface {
	ProssesMessage(message []byte) []byte
}
