package isoprepaid

import "fmt"

// AdditionalPrivateData : additional private data
type AdditionalPrivateData struct {
	SwitcherID             string `json:"switcherId,omitempty"`
	MeterSerialNumber      string `json:"meterSerialNumber,omitempty"`
	SubscriberID           string `json:"subscriberId,omitempty"`
	Flag                   string `json:"flag,omitempty"`
	PLNReferenceNumber     string `json:"plnReferenceNumber,omitempty"`
	GSPReferenceNumber     string `json:"gspReferenceNumber,omitempty"`
	SubscriberName         string `json:"subscriberName,omitempty"`
	SubscriberSegmentation string `json:"subscriberSegmentation,omitempty"`
	PowerConsumingCategory string `json:"powerConsumingCategory,omitempty"`
	MinorUnitOfAdminCharge string `json:"minorUnitOfAdminCharge,omitempty"`
	AdminCharge            string `json:"adminCharge,omitempty"`
}

// FormatInquiryString : build postpaid additional private data for inquiry request
func FormatInquiryString(data AdditionalPrivateData) (message string) {
	message = fmt.Sprintf("%07s%011s%012s%01s", data.SwitcherID, data.MeterSerialNumber, data.SubscriberID, data.Flag)

	return
}

// FormatDataString : format data string for response
func FormatDataString(data AdditionalPrivateData) (message string) {

	message = fmt.Sprintf("%07s%011s%012s%01s%-32s%-32s%-25s%-4s%09s%01s%010s",
		data.SwitcherID, data.MeterSerialNumber, data.SubscriberID, data.Flag, data.PLNReferenceNumber,
		data.GSPReferenceNumber, data.SubscriberName, data.SubscriberSegmentation, data.PowerConsumingCategory,
		data.MinorUnitOfAdminCharge, data.AdminCharge)

	return
}
