// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// InvoicePaymentMethodSpecificInput represents class InvoicePaymentMethodSpecificInput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_InvoicePaymentMethodSpecificInput
type InvoicePaymentMethodSpecificInput struct {
	AdditionalReference *string `json:"additionalReference,omitempty"`
	PaymentProductID    *int32  `json:"paymentProductId,omitempty"`
}

// NewInvoicePaymentMethodSpecificInput constructs a new InvoicePaymentMethodSpecificInput
func NewInvoicePaymentMethodSpecificInput() *InvoicePaymentMethodSpecificInput {
	return &InvoicePaymentMethodSpecificInput{}
}
