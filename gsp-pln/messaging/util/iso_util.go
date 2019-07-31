package util

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

// GetIsoCurrentTime : get iso current time
func GetIsoCurrentTime() string {
	t := time.Now()
	return fmt.Sprintf("%d%02d%02d%02d%02d%02d", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
}

// GetIsoBankCodeFormat : Standard iso 8583 for bank code
func GetIsoBankCodeFormat(bankCode string) string {
	return fmt.Sprintf("%07s", bankCode)
}

// GetIsoStanFormat : Standard System trace audit number for iso8583
func GetIsoStanFormat(stan string) string {
	return fmt.Sprintf("%012s", stan)
}

// GetIsoTerminalIDFormat : Standard terminal id format for iso8583
func GetIsoTerminalIDFormat(terminalID string) string {
	return fmt.Sprintf("%016s", terminalID)
}

// EncapsulateBytes : to encapsulate length of isobytes with length
func EncapsulateBytes(packetIso []byte) []byte {

	packetIsoLength := Bcd([]byte(fmt.Sprintf("%04x", len(string(packetIso))+2)))
	encapsulatedWithLength := append(packetIsoLength, packetIso...)

	return encapsulatedWithLength
}

// RandStringBytes : generate random string 5 characters
func RandStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// Bcd : Byte -> Binary Coded Decymal
func Bcd(data []byte) []byte {
	out := make([]byte, len(data)/2+1)
	n, err := hex.Decode(out, data)
	if err != nil {
		panic(err.Error())
	}
	return out[:n]
}
