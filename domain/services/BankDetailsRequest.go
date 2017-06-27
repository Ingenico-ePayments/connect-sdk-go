// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package services

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// BankDetailsRequest represents class BankDetailsRequest
type BankDetailsRequest struct {
	BankAccountBban *definitions.BankAccountBban `json:"bankAccountBban,omitempty"`
	BankAccountIban *definitions.BankAccountIban `json:"bankAccountIban,omitempty"`
}

// NewBankDetailsRequest constructs a new BankDetailsRequest
func NewBankDetailsRequest() *BankDetailsRequest {
	return &BankDetailsRequest{}
}
