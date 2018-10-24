// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// AbstractEInvoicePaymentMethodSpecificInput represents class AbstractEInvoicePaymentMethodSpecificInput
type AbstractEInvoicePaymentMethodSpecificInput struct {
	PaymentProductID *int32 `json:"paymentProductId,omitempty"`
	RequiresApproval *bool  `json:"requiresApproval,omitempty"`
}

// NewAbstractEInvoicePaymentMethodSpecificInput constructs a new AbstractEInvoicePaymentMethodSpecificInput
func NewAbstractEInvoicePaymentMethodSpecificInput() *AbstractEInvoicePaymentMethodSpecificInput {
	return &AbstractEInvoicePaymentMethodSpecificInput{}
}
