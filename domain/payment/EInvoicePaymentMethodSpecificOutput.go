// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// EInvoicePaymentMethodSpecificOutput represents class EInvoicePaymentMethodSpecificOutput
type EInvoicePaymentMethodSpecificOutput struct {
	FraudResults                     *definitions.FraudResults                 `json:"fraudResults,omitempty"`
	PaymentProduct9000SpecificOutput *EInvoicePaymentProduct9000SpecificOutput `json:"paymentProduct9000SpecificOutput,omitempty"`
	PaymentProductID                 *int32                                    `json:"paymentProductId,omitempty"`
}

// NewEInvoicePaymentMethodSpecificOutput constructs a new EInvoicePaymentMethodSpecificOutput
func NewEInvoicePaymentMethodSpecificOutput() *EInvoicePaymentMethodSpecificOutput {
	return &EInvoicePaymentMethodSpecificOutput{}
}
