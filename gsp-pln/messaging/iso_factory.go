package messaging

import (
	"strings"

	log "gitlab.com/kasku/kasku-2pay/2pay-billerpayment/log"

	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/config"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/isonetman"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/isonontaglis"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/isopostpaid"
	"gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/isoprepaid"
)

// GetTypeMessage : get iso message based on message type
func GetTypeMessage(messageType string) BuildIso {
	switch messageType {
	case config.Get().Mti.Netman.Request, config.Get().Mti.Netman.Response:
		log.Get().Println("IsoFactory[GetTypeMessage] netman : ", messageType)
		return &isonetman.Netman{}
	case strings.Join([]string{config.Get().Mti.Inquiry.Request, config.Get().Gsp.Nontaglis.Pan}, ""),
		strings.Join([]string{config.Get().Mti.Inquiry.Response, config.Get().Gsp.Nontaglis.Pan}, ""):
		return &isonontaglis.IsoInquiry{}
	case strings.Join([]string{config.Get().Mti.Payment.Request, config.Get().Gsp.Nontaglis.Pan}, ""),
		strings.Join([]string{config.Get().Mti.Payment.Response, config.Get().Gsp.Nontaglis.Pan}, ""):
		return &isonontaglis.IsoPayment{}
	case strings.Join([]string{config.Get().Mti.Reversal.Request, config.Get().Gsp.Nontaglis.Pan}, ""),
		strings.Join([]string{config.Get().Mti.Reversal.Response, config.Get().Gsp.Nontaglis.Pan}, ""),
		strings.Join([]string{config.Get().Mti.Reversal.Repeat.Request, config.Get().Gsp.Nontaglis.Pan}, ""),
		strings.Join([]string{config.Get().Mti.Reversal.Repeat.Response, config.Get().Gsp.Nontaglis.Pan}, ""):
		return &isonontaglis.IsoReversal{}
	case strings.Join([]string{config.Get().Mti.Inquiry.Request, config.Get().Gsp.Postpaid.Pan}, ""),
		strings.Join([]string{config.Get().Mti.Inquiry.Response, config.Get().Gsp.Postpaid.Pan}, ""):
		return &isopostpaid.IsoInquiry{}
	case strings.Join([]string{config.Get().Mti.Payment.Request, config.Get().Gsp.Postpaid.Pan}, ""),
		strings.Join([]string{config.Get().Mti.Payment.Response, config.Get().Gsp.Postpaid.Pan}, ""):
		return &isopostpaid.IsoPayment{}
	case strings.Join([]string{config.Get().Mti.Reversal.Request, config.Get().Gsp.Postpaid.Pan}, ""),
		strings.Join([]string{config.Get().Mti.Reversal.Response, config.Get().Gsp.Postpaid.Pan}, ""),
		strings.Join([]string{config.Get().Mti.Reversal.Repeat.Request, config.Get().Gsp.Postpaid.Pan}, ""),
		strings.Join([]string{config.Get().Mti.Reversal.Repeat.Response, config.Get().Gsp.Postpaid.Pan}, ""):
		return &isopostpaid.IsoReversal{}
	case strings.Join([]string{config.Get().Mti.Inquiry.Request, config.Get().Gsp.Prepaid.Pan}, ""),
		strings.Join([]string{config.Get().Mti.Inquiry.Response, config.Get().Gsp.Prepaid.Pan}, ""):
		return &isoprepaid.IsoInquiry{}
	case strings.Join([]string{config.Get().Mti.Payment.Request, config.Get().Gsp.Prepaid.Pan}, ""),
		strings.Join([]string{config.Get().Mti.Payment.Response, config.Get().Gsp.Prepaid.Pan}, ""):
		return &isoprepaid.IsoPurchase{}
	case strings.Join([]string{config.Get().Mti.Advice.Request, config.Get().Gsp.Prepaid.Pan}, ""),
		strings.Join([]string{config.Get().Mti.Advice.Response, config.Get().Gsp.Prepaid.Pan}, ""),
		strings.Join([]string{config.Get().Mti.Advice.Repeat.Request, config.Get().Gsp.Prepaid.Pan}, ""),
		strings.Join([]string{config.Get().Mti.Advice.Repeat.Response, config.Get().Gsp.Prepaid.Pan}, ""):
		return &isoprepaid.IsoPurchase{}
	default:
		return nil
	}
}
