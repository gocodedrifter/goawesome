package isoprepaid

import "fmt"

// AdditionalPrivateData : additional private data
type AdditionalPrivateData struct {
	SwitcherID                             string `json:"switcherId,omitempty"`
	MeterSerialNumber                      string `json:"meterSerialNumber,omitempty"`
	SubscriberID                           string `json:"subscriberId,omitempty"`
	Flag                                   string `json:"flag,omitempty"`
	PLNReferenceNumber                     string `json:"plnReferenceNumber,omitempty"`
	GSPReferenceNumber                     string `json:"gspReferenceNumber,omitempty"`
	VendingReceiptNumber                   string `json:"vendingReceiptNumber,omitempty"`
	SubscriberName                         string `json:"subscriberName,omitempty"`
	SubscriberSegmentation                 string `json:"subscriberSegmentation,omitempty"`
	PowerConsumingCategory                 string `json:"powerConsumingCategory,omitempty"`
	BuyingOptions                          string `json:"buyingOptions,omitempty"`
	MinorUnitOfAdminCharge                 string `json:"minorUnitOfAdminCharge,omitempty"`
	AdminCharge                            string `json:"adminCharge,omitempty"`
	MinorUnitOfStampDuty                   string `json:"minorUnitOfStampDuty,omitempty"`
	StampDuty                              string `json:"stampDuty,omitempty"`
	MinorUnitOfValueAddedTax               string `json:"minorUnitOfValueAddedTax,omitempty"`
	ValueAddedTax                          string `json:"valueAddedTax,omitempty"`
	MinorUnitOfPublicLightingTax           string `json:"minorUnitOfPublicLightingTax,omitempty"`
	PublicLightingTax                      string `json:"publicLightingTax,omitempty"`
	MinorUnitOfCustomerPayablesInstallment string `json:"minorUnitOfCustomerPayableInstallment,omitempty"`
	CustomerPayablesInstallment            string `json:"customerPayableInstallment,omitempty"`
	MinorUnitOfPowerPurchase               string `json:"minorUnitOfPowerPurchase,omitempty"`
	PowerPurchase                          string `json:"powerPurchase,omitempty"`
	MinorUnitOfPurchasedKwhUnit            string `json:"minorUnitOfPurchasedKwhUnit,omitempty"`
	PurchasedKwhUnit                       string `json:"purchasedKwhUnit,omitempty"`
	TokenNumber                            string `json:"tokenNumber,omitempty"`
}

// FormatInquiryString : build postpaid additional private data for inquiry request
func FormatInquiryString(data *AdditionalPrivateData) (message string) {
	message = fmt.Sprintf("%07s%011s%012s%01s", data.SwitcherID, data.MeterSerialNumber, data.SubscriberID, data.Flag)

	return
}

// FormatPurchaseReq : format purchase request for prepaid
func FormatPurchaseReq(data *AdditionalPrivateData) (message string) {
	message = fmt.Sprintf("%07s%011s%012s%01s%-32s%-32s%-25s%-4s%09s%01s%010s%s",
		data.SwitcherID, data.MeterSerialNumber, data.SubscriberID, data.Flag, data.PLNReferenceNumber,
		data.GSPReferenceNumber, data.SubscriberName, data.SubscriberSegmentation, data.PowerConsumingCategory,
		data.MinorUnitOfAdminCharge, data.AdminCharge, data.BuyingOptions)

	return
}

// FormatDataString : format data string for response
func FormatDataString(data *AdditionalPrivateData) (message string) {

	message = fmt.Sprintf("%07s%011s%012s%01s%-32s%-32s%-25s%-4s%09s%01s%010s",
		data.SwitcherID, data.MeterSerialNumber, data.SubscriberID, data.Flag, data.PLNReferenceNumber,
		data.GSPReferenceNumber, data.SubscriberName, data.SubscriberSegmentation, data.PowerConsumingCategory,
		data.MinorUnitOfAdminCharge, data.AdminCharge)

	return
}

// BuildInquiry : build inquiry
func BuildInquiry(message string) (data AdditionalPrivateData) {
	data.SwitcherID = message[:7]
	data.MeterSerialNumber = message[7:18]
	data.SubscriberID = message[18:30]
	data.Flag = message[30:31]
	return
}

// BuildDataResponse : parse prepaid for additional private data for inquiry response
func BuildInquiryResponse(message string) (data AdditionalPrivateData) {
	data.SwitcherID = message[:7]
	data.MeterSerialNumber = message[7:18]
	data.SubscriberID = message[18:30]
	data.Flag = message[30:31]
	data.PLNReferenceNumber = message[31:63]
	data.GSPReferenceNumber = message[63:95]
	data.SubscriberName = message[95:120]
	data.SubscriberSegmentation = message[120:124]
	data.PowerConsumingCategory = message[124:133]
	data.MinorUnitOfAdminCharge = message[133:134]
	data.AdminCharge = message[134:144]

	return
}

// BuildPurchaseResponse : parse prepaid for additional private data for purchase response
func BuildPurchaseResponse(message string) (data AdditionalPrivateData) {

	data.SwitcherID = message[:7]
	data.MeterSerialNumber = message[7:18]
	data.SubscriberID = message[18:30]
	data.Flag = message[30:31]
	data.PLNReferenceNumber = message[31:63]
	data.GSPReferenceNumber = message[63:95]
	data.VendingReceiptNumber = message[95:103]
	data.SubscriberName = message[103:128]
	data.SubscriberSegmentation = message[128:132]
	data.PowerConsumingCategory = message[132:141]
	data.BuyingOptions = message[141:142]
	data.MinorUnitOfAdminCharge = message[142:143]
	data.AdminCharge = message[143:153]
	data.MinorUnitOfStampDuty = message[153:154]
	data.StampDuty = message[154:164]
	data.MinorUnitOfValueAddedTax = message[164:165]
	data.ValueAddedTax = message[165:175]
	data.MinorUnitOfPublicLightingTax = message[175:176]
	data.PublicLightingTax = message[176:186]
	data.MinorUnitOfCustomerPayablesInstallment = message[186:187]
	data.CustomerPayablesInstallment = message[187:197]
	data.MinorUnitOfPowerPurchase = message[197:198]
	data.PowerPurchase = message[198:210]
	data.MinorUnitOfPurchasedKwhUnit = message[210:211]
	data.PurchasedKwhUnit = message[211:221]
	data.TokenNumber = message[221:]

	return
}
