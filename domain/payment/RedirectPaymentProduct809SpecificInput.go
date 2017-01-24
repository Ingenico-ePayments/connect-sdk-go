// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// RedirectPaymentProduct809SpecificInput represents class RedirectPaymentProduct809SpecificInput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_RedirectPaymentProduct809SpecificInput
type RedirectPaymentProduct809SpecificInput struct {
	ExpirationPeriod *string `json:"expirationPeriod,omitempty"`
	IssuerID         *string `json:"issuerId,omitempty"`
}

// NewRedirectPaymentProduct809SpecificInput constructs a new RedirectPaymentProduct809SpecificInput
func NewRedirectPaymentProduct809SpecificInput() *RedirectPaymentProduct809SpecificInput {
	return &RedirectPaymentProduct809SpecificInput{}
}
