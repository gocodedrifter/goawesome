package isopostpaid

import "fmt"

// BillDetail : bill detail for postpaid
type BillDetail struct {
	BillPeriod            string `json:"billPeriod,omitempty"`
	DueDate               string `json:"dueDate,omitempty"`
	MeterReadDate         string `json:"meterReadDate,omitempty"`
	TotalElectricityBill  string `json:"totalElectricityBill,omitempty"`
	Incentive             string `json:"incentive,omitempty"`
	ValueAddedTax         string `json:"valueAddedTax,omitempty"`
	PenaltyFee            string `json:"penaltyFee,omitempty"`
	PreviousMeterReading1 string `json:"previousMeterReading1,omitempty"`
	CurrentMeterReading1  string `json:"currentMeterReading1,omitempty"`
	PreviousMeterReading2 string `json:"previousMeterReading2,omitempty"`
	CurrentMeterReading2  string `json:"currentMeterReading2,omitempty"`
	PreviousMeterReading3 string `json:"previousMeterReading3,omitempty"`
	CurrentMeterReading3  string `json:"currentMeterReading3,omitempty"`
}

// ParseBillDetail : parse message to bill detail
func ParseBillDetail(message string) (billDetail BillDetail) {
	billDetail.BillPeriod = message[:6]
	billDetail.DueDate = message[6:14]
	billDetail.MeterReadDate = message[14:22]
	billDetail.TotalElectricityBill = message[22:34]
	billDetail.Incentive = message[34:45]
	billDetail.ValueAddedTax = message[45:55]
	billDetail.PenaltyFee = message[55:67]
	billDetail.PreviousMeterReading1 = message[67:75]
	billDetail.CurrentMeterReading1 = message[75:83]
	billDetail.PreviousMeterReading2 = message[83:91]
	billDetail.CurrentMeterReading2 = message[91:99]
	billDetail.PreviousMeterReading3 = message[99:107]
	billDetail.CurrentMeterReading3 = message[107:115]

	return
}

// FormatBillString : format bill string for inquiry response
func FormatBillString(billDetail BillDetail) (message string) {
	message = fmt.Sprintf("%06s%08s%08s%012s%011s%010s%012s%08s%08s%08s%08s%08s%08s", billDetail.BillPeriod, billDetail.DueDate,
		billDetail.MeterReadDate, billDetail.TotalElectricityBill, billDetail.Incentive, billDetail.ValueAddedTax, billDetail.PenaltyFee,
		billDetail.PreviousMeterReading1, billDetail.CurrentMeterReading1, billDetail.PreviousMeterReading2, billDetail.CurrentMeterReading2,
		billDetail.PreviousMeterReading3, billDetail.CurrentMeterReading3)
	return

}
