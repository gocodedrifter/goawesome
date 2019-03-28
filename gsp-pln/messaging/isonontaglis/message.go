package isonontaglis

import "gitlab.com/kasku/kasku-2pay/2pay-billerpayment/gsp-pln/messaging/basic"

// Message : message for iso non tagihan listrik
type Message struct {
	Mti                      string                  `json:"mti,omitempty"`
	PrimaryAccountNumber     string                  `json:"primaryAccountNumber,omitempty"`
	TransactionAmount        basic.TransactionAmount `json:"transactionAmount,omitempty"`
	Stan                     string                  `json:"stan,omitempty"`
	DateTimeLocalTransaction string                  `json:"dateTimeLocalTransaction,omitempty"`
	SettlementDate           string                  `json:"settlementDate,omitempty"`
	MerchantCategoryCode     string                  `json:"merchantCategoryCode,omitempty"`
	BankCode                 string                  `json:"bankCode,omitempty"`
	PartnerCentralID         string                  `json:"partnerCentralId,omitempty"`
	ResponseCode             string                  `json:"responseCode,omitempty"`
	TerminalID               string                  `json:"terminalId,omitempty"`
	AdditionalPrivateData    AdditionalPrivateData   `json:"additionalPrivateData,omitempty"`
	OriginalData             OriginalData            `json:"originalData,omitempty"`
	AdditionalPrivateData2   AdditionalPrivateData2  `json:"additionalPrivateData2,omitempty"`
	AdditionalPrivateData3   AdditionalPrivateData3  `json:"additionalPrivateData3,omitempty"`
}
