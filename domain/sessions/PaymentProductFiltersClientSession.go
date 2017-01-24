// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package sessions

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// PaymentProductFiltersClientSession represents class PaymentProductFiltersClientSession
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentProductFiltersClientSession
type PaymentProductFiltersClientSession struct {
	Exclude    *definitions.PaymentProductFilter `json:"exclude,omitempty"`
	RestrictTo *definitions.PaymentProductFilter `json:"restrictTo,omitempty"`
}

// NewPaymentProductFiltersClientSession constructs a new PaymentProductFiltersClientSession
func NewPaymentProductFiltersClientSession() *PaymentProductFiltersClientSession {
	return &PaymentProductFiltersClientSession{}
}
