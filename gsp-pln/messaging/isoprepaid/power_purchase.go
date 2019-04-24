package isoprepaid

// PowerPurchase : power purchase
type PowerPurchase struct {
	Power string `json:"power,omitempty"`
}

// BuildPower : parse message to build power purchase
func BuildPower(message string) (pp PowerPurchase) {
	pp.Power = message

	return
}
