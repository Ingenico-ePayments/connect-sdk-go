// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// AbstractPaymentMethodSpecificOutput represents class AbstractPaymentMethodSpecificOutput
type AbstractPaymentMethodSpecificOutput struct {
	PaymentProductID *int32 `json:"paymentProductId,omitempty"`
}

// NewAbstractPaymentMethodSpecificOutput constructs a new AbstractPaymentMethodSpecificOutput
func NewAbstractPaymentMethodSpecificOutput() *AbstractPaymentMethodSpecificOutput {
	return &AbstractPaymentMethodSpecificOutput{}
}
