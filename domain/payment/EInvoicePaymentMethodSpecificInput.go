// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// EInvoicePaymentMethodSpecificInput represents class EInvoicePaymentMethodSpecificInput
type EInvoicePaymentMethodSpecificInput struct {
	AcceptedTermsAndConditions      *bool                                    `json:"acceptedTermsAndConditions,omitempty"`
	PaymentProduct9000SpecificInput *EInvoicePaymentProduct9000SpecificInput `json:"paymentProduct9000SpecificInput,omitempty"`
	PaymentProductID                *int32                                   `json:"paymentProductId,omitempty"`
	RequiresApproval                *bool                                    `json:"requiresApproval,omitempty"`
}

// NewEInvoicePaymentMethodSpecificInput constructs a new EInvoicePaymentMethodSpecificInput
func NewEInvoicePaymentMethodSpecificInput() *EInvoicePaymentMethodSpecificInput {
	return &EInvoicePaymentMethodSpecificInput{}
}
