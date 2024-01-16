// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// RedirectPaymentProduct840SpecificInput represents class RedirectPaymentProduct840SpecificInput
type RedirectPaymentProduct840SpecificInput struct {
	AddressSelectionAtPayPal *bool   `json:"addressSelectionAtPayPal,omitempty"`
	// Deprecated: Use Order.references.descriptor instead
	Custom                   *string `json:"custom,omitempty"`
	IsShortcut               *bool   `json:"isShortcut,omitempty"`
}

// NewRedirectPaymentProduct840SpecificInput constructs a new RedirectPaymentProduct840SpecificInput
func NewRedirectPaymentProduct840SpecificInput() *RedirectPaymentProduct840SpecificInput {
	return &RedirectPaymentProduct840SpecificInput{}
}
