// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// AccountFundingRecipient represents class AccountFundingRecipient
type AccountFundingRecipient struct {
	AccountNumber     *string              `json:"accountNumber,omitempty"`
	AccountNumberType *string              `json:"accountNumberType,omitempty"`
	Address           *definitions.Address `json:"address,omitempty"`
	DateOfBirth       *string              `json:"dateOfBirth,omitempty"`
	Name              *AfrName             `json:"name,omitempty"`
	PartialPan        *string              `json:"partialPan,omitempty"`
}

// NewAccountFundingRecipient constructs a new AccountFundingRecipient
func NewAccountFundingRecipient() *AccountFundingRecipient {
	return &AccountFundingRecipient{}
}
