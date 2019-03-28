package basic

import (
	"github.com/Ayvan/iso8583"
)

// Iso8583Format : Iso 8583:2003 format for transaction
type Iso8583Format struct {
	PrimaryAccountNumber     *iso8583.Llvar        `field:"2" length:"5" encode:"ascii"`
	TransactionAmount        *iso8583.Alphanumeric `field:"4" length:"16" encode:"ascii"`
	Stan                     *iso8583.Alphanumeric `field:"11" length:"12" encode:"ascii"`
	DateTimeLocalTransaction *iso8583.Alphanumeric `field:"12" length:"14" encode:"ascii"`
	SettlementDate           *iso8583.Alphanumeric `field:"15" length:"8" encode:"ascii"`
	MerchantCategoryCode     *iso8583.Alphanumeric `field:"26" length:"4" encode:"ascii"`
	BankCode                 *iso8583.Llvar        `field:"32" length:"7" encode:"ascii"`
	PartnerCentralID         *iso8583.Llvar        `field:"33" length:"7" encode:"ascii"`
	ResponseCode             *iso8583.Alphanumeric `field:"39" length:"4" encode:"ascii"`
	ActionCode               *iso8583.Alphanumeric `field:"40" length:"3" encode:"ascii"`
	TerminalID               *iso8583.Alphanumeric `field:"41" length:"16" encode:"ascii"`
	AdditionalPrivateData    *iso8583.Lllvar       `field:"48" length:"999" encode:"ascii"`
	OriginalData             *iso8583.Lllvar       `field:"56" length:"999" encode:"ascii"`
	AdditionalPrivateData2   *iso8583.Lllvar       `field:"61" length:"999" encode:"ascii"`
	AdditionalPrivateData3   *iso8583.Lllvar       `field:"62" length:"999" encode:"ascii"`
	InfoText                 *iso8583.Lllvar       `field:"63" length:"999" encode:"ascii"`
}

// DecodeIsoMessage : decode iso bit message to iso8583 format
func DecodeIsoMessage(message []byte) *Iso8583Format {

	iso := iso8583.NewMessageExtended("", iso8583.ASCII, false, true,
		&Iso8583Format{
			PrimaryAccountNumber:     iso8583.NewLlvar([]byte("")),
			TransactionAmount:        iso8583.NewAlphanumeric(""),
			Stan:                     iso8583.NewAlphanumeric(""),
			DateTimeLocalTransaction: iso8583.NewAlphanumeric(""),
			SettlementDate:           iso8583.NewAlphanumeric(""),
			MerchantCategoryCode:     iso8583.NewAlphanumeric(""),
			BankCode:                 iso8583.NewLlvar([]byte("")),
			PartnerCentralID:         iso8583.NewLlvar([]byte("")),
			ResponseCode:             iso8583.NewAlphanumeric(""),
			ActionCode:               iso8583.NewAlphanumeric(""),
			TerminalID:               iso8583.NewAlphanumeric(""),
			AdditionalPrivateData:    iso8583.NewLllvar([]byte("")),
			OriginalData:             iso8583.NewLllvar([]byte("")),
			AdditionalPrivateData2:   iso8583.NewLllvar([]byte("")),
			AdditionalPrivateData3:   iso8583.NewLllvar([]byte("")),
			InfoText:                 iso8583.NewLllvar([]byte("")),
		})

	if err := iso.Load(message[2:]); err != nil {
		panic(err.Error())
	}

	resultFields := iso.Data.(*Iso8583Format)
	return resultFields
}
