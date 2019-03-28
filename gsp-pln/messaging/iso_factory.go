package messaging

import (
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/isonetman"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/isonontaglis"
)

// GetTypeMessage : get iso message based on message type
func GetTypeMessage(messageType string) BuildIso {
	switch messageType {
	case "2810":
		return &isonetman.Netman{}
	case "210099504", "211099504":
		return &isonontaglis.IsoInquiry{}
	case "220099504", "221099504":
		return &isonontaglis.IsoPayment{}
	case "240099504", "240199504", "241099504", "241199504":
		return &isonontaglis.IsoReversal{}
	default:
		return nil
	}
}
