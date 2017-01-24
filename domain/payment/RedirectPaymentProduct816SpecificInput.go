// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// RedirectPaymentProduct816SpecificInput represents class RedirectPaymentProduct816SpecificInput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_RedirectPaymentProduct816SpecificInput
type RedirectPaymentProduct816SpecificInput struct {
	BankAccountIban *definitions.BankAccountIban `json:"bankAccountIban,omitempty"`
}

// NewRedirectPaymentProduct816SpecificInput constructs a new RedirectPaymentProduct816SpecificInput
func NewRedirectPaymentProduct816SpecificInput() *RedirectPaymentProduct816SpecificInput {
	return &RedirectPaymentProduct816SpecificInput{}
}
