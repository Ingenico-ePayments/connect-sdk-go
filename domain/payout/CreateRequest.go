// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payout

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// CreateRequest represents class CreatePayoutRequest
type CreateRequest struct {
	AmountOfMoney                         *definitions.AmountOfMoney             `json:"amountOfMoney,omitempty"`
	// Deprecated: Use bankTransferPayoutMethodSpecificInput.bankAccountBban instead
	BankAccountBban                       *definitions.BankAccountBban           `json:"bankAccountBban,omitempty"`
	// Deprecated: Use bankTransferPayoutMethodSpecificInput.bankAccountIban instead
	BankAccountIban                       *definitions.BankAccountIban           `json:"bankAccountIban,omitempty"`
	BankTransferPayoutMethodSpecificInput *BankTransferPayoutMethodSpecificInput `json:"bankTransferPayoutMethodSpecificInput,omitempty"`
	CardPayoutMethodSpecificInput         *CardPayoutMethodSpecificInput         `json:"cardPayoutMethodSpecificInput,omitempty"`
	// Deprecated: Use bankTransferPayoutMethodSpecificInput.customer instead
	Customer                              *Customer                              `json:"customer,omitempty"`
	// Deprecated: Use bankTransferPayoutMethodSpecificInput.payoutDate instead
	PayoutDate                            *string                                `json:"payoutDate,omitempty"`
	// Deprecated: Use bankTransferPayoutMethodSpecificInput.payoutText instead
	PayoutText                            *string                                `json:"payoutText,omitempty"`
	References                            *References                            `json:"references,omitempty"`
	// Deprecated: Use bankTransferPayoutMethodSpecificInput.swiftCode instead
	SwiftCode                             *string                                `json:"swiftCode,omitempty"`
}

// NewCreateRequest constructs a new CreateRequest
func NewCreateRequest() *CreateRequest {
	return &CreateRequest{}
}
