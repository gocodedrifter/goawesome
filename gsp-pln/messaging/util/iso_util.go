package util

import (
	"fmt"
	"time"
)

// GetIso8583CurrentTime : get iso current time
func GetIso8583CurrentTime() string {
	t := time.Now()
	return fmt.Sprintf("%d%02d%02d%02d%02d%02d", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
}

// GetIso8583BankCodeFormat : Standard iso 8583 for bank code
func GetIso8583BankCodeFormat(bankCode string) string {
	return fmt.Sprintf("%07s", bankCode)
}

// GetIso8583StanFormat : Standard System trace audit number for iso8583
func GetIso8583StanFormat(stan string) string {
	return fmt.Sprintf("%012s", stan)
}

// GetIso8583TerminalIDFormat : Standard terminal id format for iso8583
func GetIso8583TerminalIDFormat(terminalID string) string {
	return fmt.Sprintf("%016s", terminalID)
}
