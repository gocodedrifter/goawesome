package basic

import "fmt"

// OriginalData : original data
type OriginalData struct {
	OriginalMti      string `json:"originalMti,omitempty"`
	OriginalStan     string `json:"originalStan,omitempty"`
	OriginalTime     string `json:"originalTime,omitempty"`
	OriginalBankCode string `json:"originalBankCode,omitempty"`
}

// FormatReversalString : format reversal string
func FormatReversalString(data *OriginalData) (message string) {
	message = fmt.Sprintf("%04s%012s%014s%07s", data.OriginalMti, data.OriginalStan, data.OriginalTime, data.OriginalBankCode)
	return
}

// BuildOriginalDataResponse : build original data
func BuildOriginalDataResponse(message string) (data OriginalData) {
	data.OriginalMti = message[:4]
	data.OriginalStan = message[4:16]
	data.OriginalTime = message[16:30]
	data.OriginalBankCode = message[30:]
	return
}
