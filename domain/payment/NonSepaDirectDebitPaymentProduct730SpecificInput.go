// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// NonSepaDirectDebitPaymentProduct730SpecificInput represents class NonSepaDirectDebitPaymentProduct730SpecificInput
type NonSepaDirectDebitPaymentProduct730SpecificInput struct {
	BankAccountBban *definitions.BankAccountBban `json:"bankAccountBban,omitempty"`
}

// NewNonSepaDirectDebitPaymentProduct730SpecificInput constructs a new NonSepaDirectDebitPaymentProduct730SpecificInput
func NewNonSepaDirectDebitPaymentProduct730SpecificInput() *NonSepaDirectDebitPaymentProduct730SpecificInput {
	return &NonSepaDirectDebitPaymentProduct730SpecificInput{}
}
