// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// RedirectPaymentProduct809SpecificInput represents class RedirectPaymentProduct809SpecificInput
type RedirectPaymentProduct809SpecificInput struct {
	// Deprecated: Use RedirectPaymentMethodSpecificInput.expirationPeriod instead
	ExpirationPeriod *string `json:"expirationPeriod,omitempty"`
	IssuerID         *string `json:"issuerId,omitempty"`
}

// NewRedirectPaymentProduct809SpecificInput constructs a new RedirectPaymentProduct809SpecificInput
func NewRedirectPaymentProduct809SpecificInput() *RedirectPaymentProduct809SpecificInput {
	return &RedirectPaymentProduct809SpecificInput{}
}
