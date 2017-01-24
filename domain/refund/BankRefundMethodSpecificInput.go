// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package refund

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// BankRefundMethodSpecificInput represents class BankRefundMethodSpecificInput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_BankRefundMethodSpecificInput
type BankRefundMethodSpecificInput struct {
	BankAccountBban *BankAccountBbanRefund       `json:"bankAccountBban,omitempty"`
	BankAccountIban *definitions.BankAccountIban `json:"bankAccountIban,omitempty"`
	CountryCode     *string                      `json:"countryCode,omitempty"`
}

// NewBankRefundMethodSpecificInput constructs a new BankRefundMethodSpecificInput
func NewBankRefundMethodSpecificInput() *BankRefundMethodSpecificInput {
	return &BankRefundMethodSpecificInput{}
}
