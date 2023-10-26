// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// AdditionalOrderInput represents class AdditionalOrderInput
type AdditionalOrderInput struct {
	AccountFundingRecipient *AccountFundingRecipient `json:"accountFundingRecipient,omitempty"`
	AirlineData             *definitions.AirlineData `json:"airlineData,omitempty"`
	Installments            *Installments            `json:"installments,omitempty"`
	// Deprecated: Use Order.shoppingCart.amountBreakdown instead
	Level3SummaryData       *Level3SummaryData       `json:"level3SummaryData,omitempty"`
	// Deprecated: No replacement
	LoanRecipient           *LoanRecipient           `json:"loanRecipient,omitempty"`
	LodgingData             *definitions.LodgingData `json:"lodgingData,omitempty"`
	// Deprecated: Use installments.numberOfInstallments instead
	NumberOfInstallments    *int64                   `json:"numberOfInstallments,omitempty"`
	OrderDate               *string                  `json:"orderDate,omitempty"`
	TypeInformation         *OrderTypeInformation    `json:"typeInformation,omitempty"`
}

// NewAdditionalOrderInput constructs a new AdditionalOrderInput
func NewAdditionalOrderInput() *AdditionalOrderInput {
	return &AdditionalOrderInput{}
}
