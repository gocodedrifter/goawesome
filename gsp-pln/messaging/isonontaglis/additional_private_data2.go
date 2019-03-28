package isonontaglis

import "fmt"

// AdditionalPrivateData2 : additional private data 2 for iso non tagihan listrik
type AdditionalPrivateData2 struct {
	MutationNumber         string `json:"mutationNumber,omitempty"`
	SubscriberSegmentation string `json:"subscriberSegmentation,omitempty"`
	PowerConsumingCategory string `json:"powerConsumingCategory,omitempty"`
	GSPReferenceNumber     string `json:"gspReferenceNumber,omitempty"`
}

// BuildData2Response : parse message to iso additional private data 2
func BuildData2Response(message string) (addPrivateData2 AdditionalPrivateData2) {
	addPrivateData2.MutationNumber = message[:32]
	addPrivateData2.SubscriberSegmentation = message[32:36]
	addPrivateData2.PowerConsumingCategory = message[36:45]
	addPrivateData2.GSPReferenceNumber = message[45:]

	return
}

// FormatData2String : format string for additional private data 2
func FormatData2String(data AdditionalPrivateData2) (message string) {
	message = fmt.Sprintf("%s%s%s", data.MutationNumber, data.SubscriberSegmentation, data.PowerConsumingCategory)
	return
}
