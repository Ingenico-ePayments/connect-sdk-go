// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// RedirectPaymentMethodSpecificOutput represents class RedirectPaymentMethodSpecificOutput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_RedirectPaymentMethodSpecificOutput
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
