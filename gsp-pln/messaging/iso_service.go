package messaging

// EncodeMessage : to encode the message
func EncodeMessage(buildIso BuildIso, message string) []byte {
	return buildIso.Encode(message)
}

// DecodeMessage : to decode the message
func DecodeMessage(buildIso BuildIso, isobyte []byte) (string, error) {
	return buildIso.Decode(isobyte)
}
