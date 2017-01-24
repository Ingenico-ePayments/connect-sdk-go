// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package definitions

// AbstractPaymentMethodSpecificInput represents class AbstractPaymentMethodSpecificInput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_AbstractPaymentMethodSpecificInput
type AbstractPaymentMethodSpecificInput struct {
	PaymentProductID *int32 `json:"paymentProductId,omitempty"`
}

// NewAbstractPaymentMethodSpecificInput constructs a new AbstractPaymentMethodSpecificInput
func NewAbstractPaymentMethodSpecificInput() *AbstractPaymentMethodSpecificInput {
	return &AbstractPaymentMethodSpecificInput{}
}
