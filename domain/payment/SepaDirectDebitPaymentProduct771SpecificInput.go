// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/mandates"

// SepaDirectDebitPaymentProduct771SpecificInput represents class SepaDirectDebitPaymentProduct771SpecificInput
type SepaDirectDebitPaymentProduct771SpecificInput struct {
	ExistingUniqueMandateReference *string                              `json:"existingUniqueMandateReference,omitempty"`
	Mandate                        *mandates.CreateMandateWithReturnURL `json:"mandate,omitempty"`
	// Deprecated: Use existingUniqueMandateReference or mandate.uniqueMandateReference instead
	MandateReference               *string                              `json:"mandateReference,omitempty"`
}

// NewSepaDirectDebitPaymentProduct771SpecificInput constructs a new SepaDirectDebitPaymentProduct771SpecificInput
func NewSepaDirectDebitPaymentProduct771SpecificInput() *SepaDirectDebitPaymentProduct771SpecificInput {
	return &SepaDirectDebitPaymentProduct771SpecificInput{}
}
