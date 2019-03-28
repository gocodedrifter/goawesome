package isonontaglis

import "fmt"

// AdditionalPrivateData : additional private data
type AdditionalPrivateData struct {
	SwitcherID                      string `json:"switcherId,omitempty"`
	RegistrationNumber              string `json:"registrationNumber,omitempty"`
	AreaCode                        string `json:"areaCode,omitempty"`
	TransactionCode                 string `json:"transactionCode,omitempty"`
	TransactionName                 string `json:"transactionName,omitempty"`
	RegistrationDate                string `json:"registrationDate,omitempty"`
	ExpirationDate                  string `json:"expirationDate,omitempty"`
	SubscriberID                    string `json:"subscriberId,omitempty"`
	SubscriberName                  string `json:"subscriberName,omitempty"`
	PLNReferenceNumber              string `json:"plnReferenceNumber,omitempty"`
	GSPReferenceNumber              string `json:"gspReferenceNumber,omitempty"`
	ServiceUnit                     string `json:"serviceUnit,omitempty"`
	ServiceUnitAddress              string `json:"serviceUnitAddress,omitempty"`
	ServiceUnitPhone                string `json:"serviceUnitPhone,omitempty"`
	TotalTransactionAmountMinorUnit string `json:"totalTransactionAmountMinorUnit,omitempty"`
	TotalTransactionAmount          string `json:"totalTransactionAmount,omitempty"`
	PLNBillMinorUnit                string `json:"plnBillMinorUnit,omitempty"`
	PLNBillValue                    string `json:"plnBillValue,omitempty"`
	AdministrationChargeMinorUnit   string `json:"administrationChargeMinorUnit,omitempty"`
	AdministrationCharge            string `json:"administrationCharge,omitempty"`
}

// FormatInquiryString : format inquiry for string request
func FormatInquiryString(data AdditionalPrivateData) (message string) {
	message = fmt.Sprintf("%07s%-32s%02s%03s", data.SwitcherID, data.RegistrationNumber, data.AreaCode, data.TransactionCode)
	return
}

// BuildResponse : parse non tagihan listrik additional private data for inquiry response
func BuildResponse(message string) (data AdditionalPrivateData) {
	data.SwitcherID = message[:7]
	data.RegistrationNumber = message[7:39]
	data.AreaCode = message[39:41]
	data.TransactionCode = message[41:44]
	data.TransactionName = message[44:69]
	data.RegistrationDate = message[69:77]
	data.ExpirationDate = message[77:85]
	data.SubscriberID = message[85:97]
	data.SubscriberName = message[97:122]
	data.PLNReferenceNumber = message[122:154]
	data.GSPReferenceNumber = message[154:186]
	data.ServiceUnit = message[186:191]
	data.ServiceUnitAddress = message[191:226]
	data.ServiceUnitPhone = message[226:241]
	data.TotalTransactionAmountMinorUnit = message[241:242]
	data.TotalTransactionAmount = message[242:259]
	data.PLNBillMinorUnit = message[259:260]
	data.PLNBillValue = message[260:277]
	data.AdministrationChargeMinorUnit = message[277:278]
	data.AdministrationCharge = message[278:288]

	return
}

// FormatDataString : format non taglis additional private data to string for payment request
func FormatDataString(data AdditionalPrivateData) (message string) {

	message = fmt.Sprintf("%07s%-32s%02s%03s%-25s%08s%08s%012s%-25s%032s%032s%05s%-35s%-15s%01s%017s%01s%017s%01s%010s",
		data.SwitcherID, data.RegistrationNumber, data.AreaCode, data.TransactionCode, data.TransactionName,
		data.RegistrationDate, data.ExpirationDate, data.SubscriberID, data.SubscriberName, data.PLNReferenceNumber,
		data.GSPReferenceNumber, data.ServiceUnit, data.ServiceUnitAddress, data.ServiceUnitPhone,
		data.TotalTransactionAmountMinorUnit, data.TotalTransactionAmount, data.PLNBillMinorUnit, data.PLNBillValue,
		data.AdministrationChargeMinorUnit, data.AdministrationCharge)

	return
}
