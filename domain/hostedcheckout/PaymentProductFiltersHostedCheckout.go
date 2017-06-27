// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package hostedcheckout

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// PaymentProductFiltersHostedCheckout represents class PaymentProductFiltersHostedCheckout
type PaymentProductFiltersHostedCheckout struct {
	Exclude    *definitions.PaymentProductFilter `json:"exclude,omitempty"`
	RestrictTo *definitions.PaymentProductFilter `json:"restrictTo,omitempty"`
	TokensOnly *bool                             `json:"tokensOnly,omitempty"`
}

// NewPaymentProductFiltersHostedCheckout constructs a new PaymentProductFiltersHostedCheckout
func NewPaymentProductFiltersHostedCheckout() *PaymentProductFiltersHostedCheckout {
	return &PaymentProductFiltersHostedCheckout{}
}
