package isonontaglis

import "fmt"

// AdditionalPrivateData3 : additional private data 3 for non tagihan listrik
type AdditionalPrivateData3 struct {
	BillComponentType        string `json:"billComponentType,omitempty"`
	BillComponentMinorUnit   string `json:"billComponentMinorUnit,omitempty"`
	BillComponentValueAmount string `json:"billComponentValueAmount,omitempty"`
}

// BuildData3Respose : parse message non taglis for additiona private data 3
func BuildData3Respose(message string) (addPrivateData3 AdditionalPrivateData3) {
	addPrivateData3.BillComponentType = message[:2]
	addPrivateData3.BillComponentMinorUnit = message[2:3]
	addPrivateData3.BillComponentValueAmount = message[3:]

	return
}

// FormatData3String : format non taglis additional private data 3
func FormatData3String(data AdditionalPrivateData3) (message string) {

	message = fmt.Sprintf("%02s%01s%017s", data.BillComponentType, data.BillComponentMinorUnit, data.BillComponentValueAmount)
	return
}
