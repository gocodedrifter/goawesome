package basic

import "fmt"

// TransactionAmount : transaction amount
type TransactionAmount struct {
	IsoCurrencyCode   string
	CurrencyMinorUnit string
	ValueAmount       string
}

// BuildTrxAmt : build transaction amount for iso 8583
func BuildTrxAmt(isoCurrencyCode, currencyMinorUnit, valueAmount string) (message string) {
	message = fmt.Sprintf("%03s%01s%012s", isoCurrencyCode, currencyMinorUnit, valueAmount)
	return
}

// ParseMessageToTrxAmt : parse message to Transaction amount
func ParseMessageToTrxAmt(message string) (transactionAmount TransactionAmount) {
	transactionAmount.IsoCurrencyCode = message[:3]
	transactionAmount.CurrencyMinorUnit = message[3:4]
	transactionAmount.ValueAmount = message[4:]

	return
}
