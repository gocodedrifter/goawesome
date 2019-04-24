package isoprepaid

import (
	"bytes"
	"fmt"
)

// AdditionalPrivateData3 : additional private data 3
type AdditionalPrivateData3 struct {
	DistributionCode    string          `json:"distributionCode,omitempty"`
	ServiceUnit         string          `json:"serviceUnit,omitempty"`
	ServiceUnitPhone    string          `json:"serviceUnitPhone,omitempty"`
	MaxKwhLimit         string          `json:"maxKwhLimit,omitempty"`
	TotalRepeat         string          `json:"totalRepeat,omitempty"`
	PowerPurchaseUnsold []PowerPurchase `json:"powerPurchaseUnsold,omitempty"`
}

// FormatData3String : format data 3 string for response
func FormatData3String(data AdditionalPrivateData3) (message string) {

	var power bytes.Buffer
	for _, v := range data.PowerPurchaseUnsold {
		power.WriteString(v.Power)
	}

	message = fmt.Sprintf("%02s%05s%-15s%05s%01s%s",
		data.DistributionCode, data.ServiceUnit, data.ServiceUnitPhone, data.MaxKwhLimit, data.TotalRepeat, power.String())

	return
}
