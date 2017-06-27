// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// RedirectPaymentMethodSpecificOutput represents class RedirectPaymentMethodSpecificOutput
type RedirectPaymentMethodSpecificOutput struct {
	BankAccountIban                 *definitions.BankAccountIban `json:"bankAccountIban,omitempty"`
	PaymentProduct836SpecificOutput *Product836SpecificOutput    `json:"paymentProduct836SpecificOutput,omitempty"`
	PaymentProduct840SpecificOutput *Product840SpecificOutput    `json:"paymentProduct840SpecificOutput,omitempty"`
	PaymentProductID                *int32                       `json:"paymentProductId,omitempty"`
}

// NewRedirectPaymentMethodSpecificOutput constructs a new RedirectPaymentMethodSpecificOutput
func NewRedirectPaymentMethodSpecificOutput() *RedirectPaymentMethodSpecificOutput {
	return &RedirectPaymentMethodSpecificOutput{}
}
