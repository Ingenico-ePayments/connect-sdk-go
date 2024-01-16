// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// NonSepaDirectDebitPaymentProduct705SpecificInput represents class NonSepaDirectDebitPaymentProduct705SpecificInput
type NonSepaDirectDebitPaymentProduct705SpecificInput struct {
	AuthorisationID *string                      `json:"authorisationId,omitempty"`
	BankAccountBban *definitions.BankAccountBban `json:"bankAccountBban,omitempty"`
	TransactionType *string                      `json:"transactionType,omitempty"`
}

// NewNonSepaDirectDebitPaymentProduct705SpecificInput constructs a new NonSepaDirectDebitPaymentProduct705SpecificInput
func NewNonSepaDirectDebitPaymentProduct705SpecificInput() *NonSepaDirectDebitPaymentProduct705SpecificInput {
	return &NonSepaDirectDebitPaymentProduct705SpecificInput{}
}
