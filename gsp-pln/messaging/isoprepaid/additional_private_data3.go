package isoprepaid

import (
	"bytes"
	"fmt"
	"strconv"

	log "gitlab.com/kasku/kasku-2pay/2pay-billerpayment/log"
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
func FormatData3String(data *AdditionalPrivateData3) (message string) {

	var power bytes.Buffer
	for _, v := range data.PowerPurchaseUnsold {
		power.WriteString(v.Power)
	}

	message = fmt.Sprintf("%02s%05s%-15s%05s%01s%s",
		data.DistributionCode, data.ServiceUnit, data.ServiceUnitPhone, data.MaxKwhLimit, data.TotalRepeat, power.String())

	return
}

// BuildData3 : parse prepaid for additional private data 3 for inquiry response
func BuildData3(message string) (data AdditionalPrivateData3) {
	data.DistributionCode = message[:2]
	data.ServiceUnit = message[2:7]
	data.ServiceUnitPhone = message[7:22]
	data.MaxKwhLimit = message[22:27]
	data.TotalRepeat = message[27:28]

	log.Get().Println("testing : ", data.TotalRepeat)
	log.Get().Println("rest value : ", message[28:])
	repeat, _ := strconv.Atoi(data.TotalRepeat)
	for powerLength := 0; powerLength < repeat; powerLength++ {
		log.Get().Println("power length : ", powerLength)
		data.PowerPurchaseUnsold = append(data.PowerPurchaseUnsold,
			BuildPower(message[(28+(powerLength*11)):(28+((powerLength+1)*11))]))
	}

	return
}
