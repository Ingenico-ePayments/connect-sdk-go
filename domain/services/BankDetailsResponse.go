// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package services

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// BankDetailsResponse represents class BankDetailsResponse
type BankDetailsResponse struct {
	BankAccountBban *definitions.BankAccountBban `json:"bankAccountBban,omitempty"`
	BankAccountIban *definitions.BankAccountIban `json:"bankAccountIban,omitempty"`
	BankData        *BankData                    `json:"bankData,omitempty"`
	Swift           *Swift                       `json:"swift,omitempty"`
}

// NewBankDetailsResponse constructs a new BankDetailsResponse
func NewBankDetailsResponse() *BankDetailsResponse {
	return &BankDetailsResponse{}
}
