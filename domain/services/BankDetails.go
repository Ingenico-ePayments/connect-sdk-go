// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package services

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// BankDetails represents class BankDetails
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_BankDetails
type BankDetails struct {
	BankAccountBban *definitions.BankAccountBban `json:"bankAccountBban,omitempty"`
	BankAccountIban *definitions.BankAccountIban `json:"bankAccountIban,omitempty"`
}

// NewBankDetails constructs a new BankDetails
func NewBankDetails() *BankDetails {
	return &BankDetails{}
}
