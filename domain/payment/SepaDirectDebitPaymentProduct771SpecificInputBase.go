// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/mandates"

// SepaDirectDebitPaymentProduct771SpecificInputBase represents class SepaDirectDebitPaymentProduct771SpecificInputBase
type SepaDirectDebitPaymentProduct771SpecificInputBase struct {
	ExistingUniqueMandateReference *string                     `json:"existingUniqueMandateReference,omitempty"`
	Mandate                        *mandates.CreateMandateBase `json:"mandate,omitempty"`
	// Deprecated: Use existingUniqueMandateReference or mandate.uniqueMandateReference instead
	MandateReference               *string                     `json:"mandateReference,omitempty"`
}

// NewSepaDirectDebitPaymentProduct771SpecificInputBase constructs a new SepaDirectDebitPaymentProduct771SpecificInputBase
func NewSepaDirectDebitPaymentProduct771SpecificInputBase() *SepaDirectDebitPaymentProduct771SpecificInputBase {
	return &SepaDirectDebitPaymentProduct771SpecificInputBase{}
}
