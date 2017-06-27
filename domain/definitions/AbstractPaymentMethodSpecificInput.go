// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

// AbstractPaymentMethodSpecificInput represents class AbstractPaymentMethodSpecificInput
type AbstractPaymentMethodSpecificInput struct {
	PaymentProductID *int32 `json:"paymentProductId,omitempty"`
}

// NewAbstractPaymentMethodSpecificInput constructs a new AbstractPaymentMethodSpecificInput
func NewAbstractPaymentMethodSpecificInput() *AbstractPaymentMethodSpecificInput {
	return &AbstractPaymentMethodSpecificInput{}
}
