// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// AbstractPaymentMethodSpecificOutput represents class AbstractPaymentMethodSpecificOutput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_AbstractPaymentMethodSpecificOutput
type AbstractPaymentMethodSpecificOutput struct {
	PaymentProductID *int32 `json:"paymentProductId,omitempty"`
}

// NewAbstractPaymentMethodSpecificOutput constructs a new AbstractPaymentMethodSpecificOutput
func NewAbstractPaymentMethodSpecificOutput() *AbstractPaymentMethodSpecificOutput {
	return &AbstractPaymentMethodSpecificOutput{}
}
