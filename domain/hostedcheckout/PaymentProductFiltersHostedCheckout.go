// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package hostedcheckout

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// PaymentProductFiltersHostedCheckout represents class PaymentProductFiltersHostedCheckout
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentProductFiltersHostedCheckout
type PaymentProductFiltersHostedCheckout struct {
	Exclude    *definitions.PaymentProductFilter `json:"exclude,omitempty"`
	RestrictTo *definitions.PaymentProductFilter `json:"restrictTo,omitempty"`
	TokensOnly *bool                             `json:"tokensOnly,omitempty"`
}

// NewPaymentProductFiltersHostedCheckout constructs a new PaymentProductFiltersHostedCheckout
func NewPaymentProductFiltersHostedCheckout() *PaymentProductFiltersHostedCheckout {
	return &PaymentProductFiltersHostedCheckout{}
}
