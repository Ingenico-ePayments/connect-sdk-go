// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// RedirectPaymentMethodSpecificInputBase represents class RedirectPaymentMethodSpecificInputBase
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_RedirectPaymentMethodSpecificInputBase
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
