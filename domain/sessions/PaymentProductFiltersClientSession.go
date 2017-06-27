// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package sessions

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// PaymentProductFiltersClientSession represents class PaymentProductFiltersClientSession
type PaymentProductFiltersClientSession struct {
	Exclude    *definitions.PaymentProductFilter `json:"exclude,omitempty"`
	RestrictTo *definitions.PaymentProductFilter `json:"restrictTo,omitempty"`
}

// NewPaymentProductFiltersClientSession constructs a new PaymentProductFiltersClientSession
func NewPaymentProductFiltersClientSession() *PaymentProductFiltersClientSession {
	return &PaymentProductFiltersClientSession{}
}
