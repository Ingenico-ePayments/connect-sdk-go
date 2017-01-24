// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// InvoicePaymentMethodSpecificOutput represents class InvoicePaymentMethodSpecificOutput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_InvoicePaymentMethodSpecificOutput
type InvoicePaymentMethodSpecificOutput struct {
	PaymentProductID *int32 `json:"paymentProductId,omitempty"`
}

// NewInvoicePaymentMethodSpecificOutput constructs a new InvoicePaymentMethodSpecificOutput
func NewInvoicePaymentMethodSpecificOutput() *InvoicePaymentMethodSpecificOutput {
	return &InvoicePaymentMethodSpecificOutput{}
}
