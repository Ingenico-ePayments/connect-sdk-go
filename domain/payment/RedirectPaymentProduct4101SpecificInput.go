// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// RedirectPaymentProduct4101SpecificInput represents class RedirectPaymentProduct4101SpecificInput
type RedirectPaymentProduct4101SpecificInput struct {
	IntegrationType *string `json:"integrationType,omitempty"`
	MerchantName    *string `json:"merchantName,omitempty"`
	TransactionNote *string `json:"transactionNote,omitempty"`
	Vpa             *string `json:"vpa,omitempty"`
}

// NewRedirectPaymentProduct4101SpecificInput constructs a new RedirectPaymentProduct4101SpecificInput
func NewRedirectPaymentProduct4101SpecificInput() *RedirectPaymentProduct4101SpecificInput {
	return &RedirectPaymentProduct4101SpecificInput{}
}
