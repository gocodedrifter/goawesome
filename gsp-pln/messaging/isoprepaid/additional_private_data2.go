package isoprepaid

import "fmt"

// AdditionalPrivateData2 : additional private data 2
type AdditionalPrivateData2 struct {
	GSPReferenceNumber string `json:"gspReferenceNumber,omitempty"`
}

// BuildData2Response : parse message to iso additional private data 2
func BuildData2Response(message string) (addPrivateData2 AdditionalPrivateData2) {
	addPrivateData2.GSPReferenceNumber = message

	return
}

// FormatData2String : format string for additional private data 2
func FormatData2String(data *AdditionalPrivateData2) (message string) {
	message = fmt.Sprintf("%s", data.GSPReferenceNumber)
	return
}
