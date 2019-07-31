package isoprepaid

import (
	"fmt"
	"log"
)

// AdditionalPrivateData : additional private data
type AdditionalPrivateData struct {
	SwitcherID                             string `bson:"switcherId" json:"switcherId,omitempty"`
	MeterSerialNumber                      string `bson:"meterSerialNumber" json:"meterSerialNumber,omitempty"`
	SubscriberID                           string `bson:"subscriberId" json:"subscriberId,omitempty"`
	Flag                                   string `bson:"flag" json:"flag,omitempty"`
	PLNReferenceNumber                     string `bson:"plnReferenceNumber" json:"plnReferenceNumber,omitempty"`
	GSPReferenceNumber                     string `bson:"gspReferenceNumber" json:"gspReferenceNumber,omitempty"`
	VendingReceiptNumber                   string `bson:"vendingReceiptNumber" json:"vendingReceiptNumber,omitempty"`
	SubscriberName                         string `bson:"subscriberName" json:"subscriberName,omitempty"`
	SubscriberSegmentation                 string `bson:"subscriberSegmentation" json:"subscriberSegmentation,omitempty"`
	PowerConsumingCategory                 string `bson:"powerConsumingCategory" json:"powerConsumingCategory,omitempty"`
	BuyingOptions                          string `bson:"buyingOptions" json:"buyingOptions,omitempty"`
	MinorUnitOfAdminCharge                 string `bson:"minorUnitOfAdminCharge" json:"minorUnitOfAdminCharge,omitempty"`
	AdminCharge                            string `bson:"adminCharge" json:"adminCharge,omitempty"`
	MinorUnitOfStampDuty                   string `bson:"minorUnitOfStampDuty" json:"minorUnitOfStampDuty,omitempty"`
	StampDuty                              string `bson:"stampDuty" json:"stampDuty,omitempty"`
	MinorUnitOfValueAddedTax               string `bson:"minorUnitOfValueAddedTax" json:"minorUnitOfValueAddedTax,omitempty"`
	ValueAddedTax                          string `bson:"valueAddedTax" json:"valueAddedTax,omitempty"`
	MinorUnitOfPublicLightingTax           string `bson:"minorUnitOfPublicLightingTax" json:"minorUnitOfPublicLightingTax,omitempty"`
	PublicLightingTax                      string `bson:"publicLightingTax" json:"publicLightingTax,omitempty"`
	MinorUnitOfCustomerPayablesInstallment string `bson:"minorUnitOfCustomerPayableInstallment" json:"minorUnitOfCustomerPayableInstallment,omitempty"`
	CustomerPayablesInstallment            string `bson:"customerPayableInstallment" json:"customerPayableInstallment,omitempty"`
	MinorUnitOfPowerPurchase               string `bson:"minorUnitOfPowerPurchase" json:"minorUnitOfPowerPurchase,omitempty"`
	PowerPurchase                          string `bson:"powerPurchase" json:"powerPurchase,omitempty"`
	MinorUnitOfPurchasedKwhUnit            string `bson:"minorUnitOfPurchasedKwhUnit" json:"minorUnitOfPurchasedKwhUnit,omitempty"`
	PurchasedKwhUnit                       string `bson:"purchasedKwhUnit" json:"purchasedKwhUnit,omitempty"`
	TokenNumber                            string `bson:"tokenNumber" json:"tokenNumber,omitempty"`
}

// FormatInqReq : build postpaid additional private data for inquiry request
func FormatInqReq(data *AdditionalPrivateData) (message string) {
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

// FormatPurchaseRes : format purchase response for prepaid
func FormatPurchaseRes(data *AdditionalPrivateData) (message string) {

	message = fmt.Sprintf("%07s%011s%012s%01s%-32s%-32s%08s%-25s%-4s%09s%01s%01s%010s%01s%010s%01s%010s%01s%010s%01s%010s%01s%012s%01s%010s%020s",
		data.SwitcherID, data.MeterSerialNumber, data.SubscriberID, data.Flag, data.PLNReferenceNumber,
		data.GSPReferenceNumber, data.VendingReceiptNumber, data.SubscriberName, data.SubscriberSegmentation,
		data.PowerConsumingCategory, data.BuyingOptions, data.MinorUnitOfAdminCharge, data.AdminCharge, data.MinorUnitOfStampDuty,
		data.StampDuty, data.MinorUnitOfValueAddedTax, data.ValueAddedTax, data.MinorUnitOfPublicLightingTax, data.PublicLightingTax,
		data.MinorUnitOfCustomerPayablesInstallment, data.CustomerPayablesInstallment, data.MinorUnitOfPowerPurchase, data.PowerPurchase,
		data.MinorUnitOfPurchasedKwhUnit, data.PurchasedKwhUnit, data.TokenNumber)

	return
}

// FormatInqRes : format data string for response
func FormatInqRes(data *AdditionalPrivateData) (message string) {

	message = fmt.Sprintf("%07s%011s%012s%01s%-32s%-32s%-25s%-4s%09s%01s%010s",
		data.SwitcherID, data.MeterSerialNumber, data.SubscriberID, data.Flag, data.PLNReferenceNumber,
		data.GSPReferenceNumber, data.SubscriberName, data.SubscriberSegmentation, data.PowerConsumingCategory,
		data.MinorUnitOfAdminCharge, data.AdminCharge)

	return
}

// BuildInquiryReq : build inquiry
func BuildInquiryReq(message string) (data AdditionalPrivateData) {
	data.SwitcherID = message[:7]
	data.MeterSerialNumber = message[7:18]
	data.SubscriberID = message[18:30]
	data.Flag = message[30:31]
	return
}

// BuildInquiryResponse : parse prepaid for additional private data for inquiry response
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

	if len(message) > 144 {
		data.BuyingOptions = message[144:145]
	}
	return
}

// BuildPurchaseResponse : parse prepaid for additional private data for purchase response
func BuildPurchaseResponse(message string) (data AdditionalPrivateData) {
	log.Println("data : ", message)
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

// BuildPurchaseRequest : parse prepaid for additional private data for purchase response
func BuildPurchaseRequest(message string) (data AdditionalPrivateData) {

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
	data.BuyingOptions = message[144:]

	return
}
