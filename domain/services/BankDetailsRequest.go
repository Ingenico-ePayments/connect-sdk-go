// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package services

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// BankDetailsRequest represents class BankDetailsRequest
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_BankDetailsRequest
type BankDetailsRequest struct {
	BankAccountBban *definitions.BankAccountBban `json:"bankAccountBban,omitempty"`
	BankAccountIban *definitions.BankAccountIban `json:"bankAccountIban,omitempty"`
}

// NewBankDetailsRequest constructs a new BankDetailsRequest
func NewBankDetailsRequest() *BankDetailsRequest {
	return &BankDetailsRequest{}
}
