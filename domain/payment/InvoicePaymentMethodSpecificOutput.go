// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// InvoicePaymentMethodSpecificOutput represents class InvoicePaymentMethodSpecificOutput
type InvoicePaymentMethodSpecificOutput struct {
	PaymentProductID *int32 `json:"paymentProductId,omitempty"`
}

// NewInvoicePaymentMethodSpecificOutput constructs a new InvoicePaymentMethodSpecificOutput
func NewInvoicePaymentMethodSpecificOutput() *InvoicePaymentMethodSpecificOutput {
	return &InvoicePaymentMethodSpecificOutput{}
}
