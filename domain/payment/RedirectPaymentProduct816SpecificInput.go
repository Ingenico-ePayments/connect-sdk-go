// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// RedirectPaymentProduct816SpecificInput represents class RedirectPaymentProduct816SpecificInput
type RedirectPaymentProduct816SpecificInput struct {
	BankAccountIban *definitions.BankAccountIban `json:"bankAccountIban,omitempty"`
}

// NewRedirectPaymentProduct816SpecificInput constructs a new RedirectPaymentProduct816SpecificInput
func NewRedirectPaymentProduct816SpecificInput() *RedirectPaymentProduct816SpecificInput {
	return &RedirectPaymentProduct816SpecificInput{}
}
