// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// InvoicePaymentMethodSpecificInput represents class InvoicePaymentMethodSpecificInput
type InvoicePaymentMethodSpecificInput struct {
	AdditionalReference *string `json:"additionalReference,omitempty"`
	PaymentProductID    *int32  `json:"paymentProductId,omitempty"`
}

// NewInvoicePaymentMethodSpecificInput constructs a new InvoicePaymentMethodSpecificInput
func NewInvoicePaymentMethodSpecificInput() *InvoicePaymentMethodSpecificInput {
	return &InvoicePaymentMethodSpecificInput{}
}
