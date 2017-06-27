// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// RedirectPaymentMethodSpecificInput represents class RedirectPaymentMethodSpecificInput
type RedirectPaymentMethodSpecificInput struct {
	ExpirationPeriod                  *int32                                  `json:"expirationPeriod,omitempty"`
	IsRecurring                       *bool                                   `json:"isRecurring,omitempty"`
	PaymentProduct809SpecificInput    *RedirectPaymentProduct809SpecificInput `json:"paymentProduct809SpecificInput,omitempty"`
	PaymentProduct816SpecificInput    *RedirectPaymentProduct816SpecificInput `json:"paymentProduct816SpecificInput,omitempty"`
	PaymentProduct882SpecificInput    *RedirectPaymentProduct882SpecificInput `json:"paymentProduct882SpecificInput,omitempty"`
	PaymentProductID                  *int32                                  `json:"paymentProductId,omitempty"`
	RecurringPaymentSequenceIndicator *string                                 `json:"recurringPaymentSequenceIndicator,omitempty"`
	ReturnURL                         *string                                 `json:"returnUrl,omitempty"`
	Token                             *string                                 `json:"token,omitempty"`
}

// NewRedirectPaymentMethodSpecificInput constructs a new RedirectPaymentMethodSpecificInput
func NewRedirectPaymentMethodSpecificInput() *RedirectPaymentMethodSpecificInput {
	return &RedirectPaymentMethodSpecificInput{}
}
