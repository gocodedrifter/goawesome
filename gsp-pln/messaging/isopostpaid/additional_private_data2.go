package isopostpaid

import "fmt"

// AdditionalPrivateData2 : additional private data 2 for postpaid
type AdditionalPrivateData2 struct {
	InquiryReferenceNumber string `json:"inquiryReferenceNumber,omitempty"`
}

// BuildData2Response : parse message to iso additional private data 2
func BuildData2Response(message string) (addPrivateData2 AdditionalPrivateData2) {
	addPrivateData2.InquiryReferenceNumber = message[:]

	return
}

// FormatData2String : format string for additional private data 2
func FormatData2String(data *AdditionalPrivateData2) (message string) {
	message = fmt.Sprintf("%s", data.InquiryReferenceNumber)
	return
}
