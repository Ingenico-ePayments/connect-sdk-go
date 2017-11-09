// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// EInvoicePaymentMethodSpecificInputBase represents class EInvoicePaymentMethodSpecificInputBase
type EInvoicePaymentMethodSpecificInputBase struct {
	PaymentProductID *int32 `json:"paymentProductId,omitempty"`
	RequiresApproval *bool  `json:"requiresApproval,omitempty"`
}

// NewEInvoicePaymentMethodSpecificInputBase constructs a new EInvoicePaymentMethodSpecificInputBase
func NewEInvoicePaymentMethodSpecificInputBase() *EInvoicePaymentMethodSpecificInputBase {
	return &EInvoicePaymentMethodSpecificInputBase{}
}
