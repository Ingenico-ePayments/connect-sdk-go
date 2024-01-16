// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package refund

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// BankRefundMethodSpecificInput represents class BankRefundMethodSpecificInput
type BankRefundMethodSpecificInput struct {
	BankAccountBban *BankAccountBbanRefund       `json:"bankAccountBban,omitempty"`
	BankAccountIban *definitions.BankAccountIban `json:"bankAccountIban,omitempty"`
	CountryCode     *string                      `json:"countryCode,omitempty"`
}

// NewBankRefundMethodSpecificInput constructs a new BankRefundMethodSpecificInput
func NewBankRefundMethodSpecificInput() *BankRefundMethodSpecificInput {
	return &BankRefundMethodSpecificInput{}
}
