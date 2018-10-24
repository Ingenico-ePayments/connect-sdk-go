// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// AbstractRedirectPaymentMethodSpecificInput represents class AbstractRedirectPaymentMethodSpecificInput
type AbstractRedirectPaymentMethodSpecificInput struct {
	ExpirationPeriod                  *int32  `json:"expirationPeriod,omitempty"`
	PaymentProductID                  *int32  `json:"paymentProductId,omitempty"`
	RecurringPaymentSequenceIndicator *string `json:"recurringPaymentSequenceIndicator,omitempty"`
	RequiresApproval                  *bool   `json:"requiresApproval,omitempty"`
	Token                             *string `json:"token,omitempty"`
	Tokenize                          *bool   `json:"tokenize,omitempty"`
}

// NewAbstractRedirectPaymentMethodSpecificInput constructs a new AbstractRedirectPaymentMethodSpecificInput
func NewAbstractRedirectPaymentMethodSpecificInput() *AbstractRedirectPaymentMethodSpecificInput {
	return &AbstractRedirectPaymentMethodSpecificInput{}
}
