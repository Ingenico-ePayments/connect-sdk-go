// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// InvoicePaymentMethodSpecificOutput represents class InvoicePaymentMethodSpecificOutput
type InvoicePaymentMethodSpecificOutput struct {
	FraudResults     *definitions.FraudResults `json:"fraudResults,omitempty"`
	PaymentProductID *int32                    `json:"paymentProductId,omitempty"`
}

// NewInvoicePaymentMethodSpecificOutput constructs a new InvoicePaymentMethodSpecificOutput
func NewInvoicePaymentMethodSpecificOutput() *InvoicePaymentMethodSpecificOutput {
	return &InvoicePaymentMethodSpecificOutput{}
}
