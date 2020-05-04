package basic

// Message : basic message for all message
type Message struct {
	Mti                      string      `bson:"mti" json:"mti,omitempty"`
	PrimaryAccountNumber     string      `bson:"primaryAccountNumber" json:"primaryAccountNumber,omitempty"`
	TransactionAmount        interface{} `bson:"transactionAmount" json:"transactionAmount,omitempty"`
	Stan                     string      `bson:"stan" json:"stan,omitempty"`
	DateTimeLocalTransaction string      `bson:"dateTimeLocalTransaction" json:"dateTimeLocalTransaction,omitempty"`
	SettlementDate           string      `bson:"settlementDate" json:"settlementDate,omitempty"`
	MerchantCategoryCode     string      `bson:"merchantCategoryCode" json:"merchantCategoryCode,omitempty"`
	BankCode                 string      `bson:"bankCode" json:"bankCode,omitempty"`
	PartnerCentralID         string      `bson:"partnerCentralId" json:"partnerCentralId,omitempty"`
	ResponseCode             string      `bson:"responseCode" json:"responseCode,omitempty"`
	TerminalID               string      `bson:"terminalId" json:"terminalId,omitempty"`
	AdditionalPrivateData    interface{} `bson:"additionalPrivateData" json:"additionalPrivateData,omitempty"`
	OriginalData             interface{} `bson:"originalData" json:"originalData,omitempty"`
	AdditionalPrivateData2   interface{} `bson:"additionalPrivateData2" json:"additionalPrivateData2,omitempty"`
	AdditionalPrivateData3   interface{} `bson:"additionalPrivateData3" json:"additionalPrivateData3,omitempty"`
	InfoText                 string      `bson:"infoText" json:"infoText,omitempty"`
	JSONFormat               string      `bson:"jsonFormat" json:"jsonFormat,omitempty"`
	Sleep                    string      `bson:"sleep" json:"sleep,omitempty"`
}
