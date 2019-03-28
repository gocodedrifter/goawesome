package isonetman

import "fmt"

// AdditionalPrivateData : additional private data
type AdditionalPrivateData struct {
	SwitcherID string `json:"switcherId,omitempty"`
}

// FormatString : format additional private data for for network management
func FormatString(data AdditionalPrivateData) string {
	return fmt.Sprintf("%07s", data.SwitcherID)
}

// BuildAdditionalPrivateData : build Additional Private Data Object
func BuildAdditionalPrivateData(message string) (obj AdditionalPrivateData) {
	obj.SwitcherID = message[:]
	return
}
