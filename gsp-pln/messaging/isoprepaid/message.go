package isoprepaid

import "gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/basic"

// Message : message for isoprepaid
type Message struct {
	basic.Message
	AdditionalPrivateData  AdditionalPrivateData  `json:"additionalPrivateData,omitempty"`
	AdditionalPrivateData3 AdditionalPrivateData3 `json:"additionalPrivateData3,omitempty"`
	InfoText               string                 `json:"infoText,omitempty"`
}
