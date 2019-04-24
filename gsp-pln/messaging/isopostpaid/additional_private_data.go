package isopostpaid

import (
	"bytes"
	"fmt"
)

// AdditionalPrivateData : Additional Private Data for postpaid
type AdditionalPrivateData struct {
	SwitcherID             string       `json:"switcherId,omitempty"`
	SubscriberID           string       `json:"subscriberId,omitempty"`
	BillStatus             string       `json:"billStatus,omitempty"`
	PaymentStatus          string       `json:"paymentStatus,omitempty"`
	TotalOutstandingBill   string       `json:"totalOutstandingBill,omitempty"`
	GSPReferenceNumber     string       `json:"gspReferenceNumber,omitempty"`
	SubscriberName         string       `json:"subscriberName,omitempty"`
	ServiceUnit            string       `json:"serviceUnit,omitempty"`
	ServiceUnitPhone       string       `json:"serviceUnitPhone,omitempty"`
	SubscriberSegmentation string       `json:"subscriberSegmentation,omitempty"`
	PowerConsumingCategory string       `json:"powerConsumingCategory,omitempty"`
	TotalAdminCharges      string       `json:"totalAdminCharges,omitempty"`
	BillDetail             []BillDetail `json:"billDetail,omitempty"`
}

// FormatInquiryString : build postpaid additional private data for inquiry request
func FormatInquiryString(data *AdditionalPrivateData) (message string) {
	message = fmt.Sprintf("%07s%012s", data.SwitcherID, data.SubscriberID)

	return
}

// FormatDataString : format data string for response
func FormatDataString(data *AdditionalPrivateData) (message string) {

	var bill bytes.Buffer
	for _, v := range data.BillDetail {
		bill.WriteString(FormatBillString(v))
	}

	message = fmt.Sprintf("%07s%012s%01s%s%02s%032s%-25s%05s%-15s%-4s%09s%09s%s",
		data.SwitcherID, data.SubscriberID, data.BillStatus, data.PaymentStatus, data.TotalOutstandingBill, data.GSPReferenceNumber, data.SubscriberName,
		data.ServiceUnit, data.ServiceUnitPhone, data.SubscriberSegmentation, data.PowerConsumingCategory,
		data.TotalAdminCharges, bill.String())

	return
}

// ParsePostpaidAdditionalPrivateDataForInquiryResponse : parse postpaid for additional private data for inquiry response
func ParsePostpaidAdditionalPrivateDataForInquiryResponse(message string) (additionalPrivateData AdditionalPrivateData) {
	additionalPrivateData.SwitcherID = message[:7]
	additionalPrivateData.SubscriberID = message[7:19]
	additionalPrivateData.BillStatus = message[19:20]
	additionalPrivateData.TotalOutstandingBill = message[20:22]
	additionalPrivateData.GSPReferenceNumber = message[22:54]
	additionalPrivateData.SubscriberName = message[54:79]
	additionalPrivateData.ServiceUnit = message[79:84]
	additionalPrivateData.ServiceUnitPhone = message[84:99]
	additionalPrivateData.SubscriberSegmentation = message[99:103]
	additionalPrivateData.PowerConsumingCategory = message[103:112]
	additionalPrivateData.TotalAdminCharges = message[112:121]

	billMessage := message[121:]
	for billLength := 0; billLength < len(billMessage)/115; billLength++ {
		additionalPrivateData.BillDetail = append(additionalPrivateData.BillDetail,
			ParseBillDetail(billMessage[(billLength*115):((billLength+1)*115)]))
	}

	return
}

// BuildInquiry : build inquiry
func BuildInquiry(message string) (data AdditionalPrivateData) {
	data.SwitcherID = message[:7]
	data.SubscriberID = message[7:19]
	return
}

// BuildDataResponse : parse postpaid for additional private data for inquiry response
func BuildDataResponse(message string) (additionalPrivateData AdditionalPrivateData) {
	additionalPrivateData.SwitcherID = message[:7]
	additionalPrivateData.SubscriberID = message[7:19]
	additionalPrivateData.BillStatus = message[19:20]
	additionalPrivateData.PaymentStatus = message[20:21]
	additionalPrivateData.TotalOutstandingBill = message[21:23]
	additionalPrivateData.GSPReferenceNumber = message[23:55]
	additionalPrivateData.SubscriberName = message[55:80]
	additionalPrivateData.ServiceUnit = message[80:85]
	additionalPrivateData.ServiceUnitPhone = message[85:100]
	additionalPrivateData.SubscriberSegmentation = message[100:104]
	additionalPrivateData.PowerConsumingCategory = message[104:113]
	additionalPrivateData.TotalAdminCharges = message[113:122]

	billMessage := message[122:]
	for billLength := 0; billLength < len(billMessage)/115; billLength++ {
		additionalPrivateData.BillDetail = append(additionalPrivateData.BillDetail,
			ParseBillDetail(billMessage[(billLength*115):((billLength+1)*115)]))
	}

	return
}
