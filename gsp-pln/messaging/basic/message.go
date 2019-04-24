package basic

// Message : basic message for all message
type Message struct {
	Mti                      string            `json:"mti,omitempty"`
	Payload                  string            `json:"payload,omitempty"`
	PrimaryAccountNumber     string            `json:"primaryAccountNumber,omitempty"`
	TransactionAmount        TransactionAmount `json:"transactionAmount,omitempty"`
	Stan                     string            `json:"stan,omitempty"`
	DateTimeLocalTransaction string            `json:"dateTimeLocalTransaction,omitempty"`
	SettlementDate           string            `json:"settlementDate,omitempty"`
	MerchantCategoryCode     string            `json:"merchantCategoryCode,omitempty"`
	BankCode                 string            `json:"bankCode,omitempty"`
	PartnerCentralID         string            `json:"partnerCentralId,omitempty"`
	ResponseCode             string            `json:"responseCode,omitempty"`
	TerminalID               string            `json:"terminalId,omitempty"`
	AdditionalPrivateData    interface{}       `json:"additionalPrivateData,omitempty"`
	OriginalData             OriginalData      `json:"originalData,omitempty"`
	AdditionalPrivateData2   interface{}       `json:"additionalPrivateData2,omitempty"`
	AdditionalPrivateData3   interface{}       `json:"additionalPrivateData3,omitempty"`
}
