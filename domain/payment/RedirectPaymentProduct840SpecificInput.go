// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// RedirectPaymentProduct840SpecificInput represents class RedirectPaymentProduct840SpecificInput
type RedirectPaymentProduct840SpecificInput struct {
	Custom     *string `json:"custom,omitempty"`
	IsShortcut *bool   `json:"isShortcut,omitempty"`
}

// NewRedirectPaymentProduct840SpecificInput constructs a new RedirectPaymentProduct840SpecificInput
func NewRedirectPaymentProduct840SpecificInput() *RedirectPaymentProduct840SpecificInput {
	return &RedirectPaymentProduct840SpecificInput{}
}
