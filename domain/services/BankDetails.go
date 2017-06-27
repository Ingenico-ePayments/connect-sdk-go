// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package services

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// BankDetails represents class BankDetails
type BankDetails struct {
	BankAccountBban *definitions.BankAccountBban `json:"bankAccountBban,omitempty"`
	BankAccountIban *definitions.BankAccountIban `json:"bankAccountIban,omitempty"`
}

// NewBankDetails constructs a new BankDetails
func NewBankDetails() *BankDetails {
	return &BankDetails{}
}
