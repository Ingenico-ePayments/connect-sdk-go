// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// RedirectPaymentMethodSpecificInputBase represents class RedirectPaymentMethodSpecificInputBase
type RedirectPaymentMethodSpecificInputBase struct {
	ExpirationPeriod                  *int32  `json:"expirationPeriod,omitempty"`
	PaymentProductID                  *int32  `json:"paymentProductId,omitempty"`
	RecurringPaymentSequenceIndicator *string `json:"recurringPaymentSequenceIndicator,omitempty"`
	Token                             *string `json:"token,omitempty"`
}

// NewRedirectPaymentMethodSpecificInputBase constructs a new RedirectPaymentMethodSpecificInputBase
func NewRedirectPaymentMethodSpecificInputBase() *RedirectPaymentMethodSpecificInputBase {
	return &RedirectPaymentMethodSpecificInputBase{}
}
