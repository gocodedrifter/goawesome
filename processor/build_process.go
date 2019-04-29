package processor

// BuildProcess : build process message
type BuildProcess interface {
	EncodeMessage(message []byte) []byte
	DecodeMessage(message []byte) string
}
