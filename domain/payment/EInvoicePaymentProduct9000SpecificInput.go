// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// EInvoicePaymentProduct9000SpecificInput represents class EInvoicePaymentProduct9000SpecificInput
type EInvoicePaymentProduct9000SpecificInput struct {
	BankAccountIban *definitions.BankAccountIban `json:"bankAccountIban,omitempty"`
	InstallmentID   *string                      `json:"installmentId,omitempty"`
}

// NewEInvoicePaymentProduct9000SpecificInput constructs a new EInvoicePaymentProduct9000SpecificInput
func NewEInvoicePaymentProduct9000SpecificInput() *EInvoicePaymentProduct9000SpecificInput {
	return &EInvoicePaymentProduct9000SpecificInput{}
}
