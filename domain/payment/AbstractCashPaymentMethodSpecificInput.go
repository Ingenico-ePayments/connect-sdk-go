// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// AbstractCashPaymentMethodSpecificInput represents class AbstractCashPaymentMethodSpecificInput
type AbstractCashPaymentMethodSpecificInput struct {
	PaymentProductID *int32 `json:"paymentProductId,omitempty"`
}

// NewAbstractCashPaymentMethodSpecificInput constructs a new AbstractCashPaymentMethodSpecificInput
func NewAbstractCashPaymentMethodSpecificInput() *AbstractCashPaymentMethodSpecificInput {
	return &AbstractCashPaymentMethodSpecificInput{}
}
