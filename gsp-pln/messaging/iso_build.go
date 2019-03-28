package messaging

// BuildIso : interface to process the ISO message
type BuildIso interface {
	Encode(message string) []byte
	Decode(isobyte []byte) (string, error)
}
