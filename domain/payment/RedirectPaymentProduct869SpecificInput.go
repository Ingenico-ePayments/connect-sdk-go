// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// RedirectPaymentProduct869SpecificInput represents class RedirectPaymentProduct869SpecificInput
type RedirectPaymentProduct869SpecificInput struct {
	IssuerID         *string `json:"issuerId,omitempty"`
	ResidentIDName   *string `json:"residentIdName,omitempty"`
	ResidentIDNumber *string `json:"residentIdNumber,omitempty"`
}

// NewRedirectPaymentProduct869SpecificInput constructs a new RedirectPaymentProduct869SpecificInput
func NewRedirectPaymentProduct869SpecificInput() *RedirectPaymentProduct869SpecificInput {
	return &RedirectPaymentProduct869SpecificInput{}
}
