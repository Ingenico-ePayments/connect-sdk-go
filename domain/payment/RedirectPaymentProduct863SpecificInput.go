// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// RedirectPaymentProduct863SpecificInput represents class RedirectPaymentProduct863SpecificInput
type RedirectPaymentProduct863SpecificInput struct {
	IntegrationType *string `json:"integrationType,omitempty"`
	OpenID          *string `json:"openId,omitempty"`
}

// NewRedirectPaymentProduct863SpecificInput constructs a new RedirectPaymentProduct863SpecificInput
func NewRedirectPaymentProduct863SpecificInput() *RedirectPaymentProduct863SpecificInput {
	return &RedirectPaymentProduct863SpecificInput{}
}
