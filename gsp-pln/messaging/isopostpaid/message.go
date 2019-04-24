package isopostpaid

import "gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/basic"

// Message : message for iso non tagihan listrik
type Message struct {
	basic.Message
	AdditionalPrivateData  interface{}        `json:"additionalPrivateData,omitempty"`
	OriginalData           basic.OriginalData `json:"originalData,omitempty"`
	AdditionalPrivateData2 interface{}        `json:"additionalPrivateData2,omitempty"`
}
