// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payout

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// CreateRequest represents class CreatePayoutRequest
type CreateRequest struct {
	// Deprecated: Moved to PayoutDetails
	AmountOfMoney                         *definitions.AmountOfMoney             `json:"amountOfMoney,omitempty"`
	// Deprecated: Moved to BankTransferPayoutMethodSpecificInput
	BankAccountBban                       *definitions.BankAccountBban           `json:"bankAccountBban,omitempty"`
	// Deprecated: Moved to BankTransferPayoutMethodSpecificInput
	BankAccountIban                       *definitions.BankAccountIban           `json:"bankAccountIban,omitempty"`
	BankTransferPayoutMethodSpecificInput *BankTransferPayoutMethodSpecificInput `json:"bankTransferPayoutMethodSpecificInput,omitempty"`
	CardPayoutMethodSpecificInput         *CardPayoutMethodSpecificInput         `json:"cardPayoutMethodSpecificInput,omitempty"`
	// Deprecated: Moved to PayoutDetails
	Customer                              *Customer                              `json:"customer,omitempty"`
	Merchant                              *Merchant                              `json:"merchant,omitempty"`
	// Deprecated: Moved to BankTransferPayoutMethodSpecificInput
	PayoutDate                            *string                                `json:"payoutDate,omitempty"`
	PayoutDetails                         *Details                               `json:"payoutDetails,omitempty"`
	// Deprecated: Moved to BankTransferPayoutMethodSpecificInput
	PayoutText                            *string                                `json:"payoutText,omitempty"`
	// Deprecated: Moved to PayoutDetails
	References                            *References                            `json:"references,omitempty"`
	// Deprecated: Moved to BankTransferPayoutMethodSpecificInput
	SwiftCode                             *string                                `json:"swiftCode,omitempty"`
}

// NewCreateRequest constructs a new CreateRequest
func NewCreateRequest() *CreateRequest {
	return &CreateRequest{}
}
