// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/mandates"

// SepaDirectDebitPaymentProduct771SpecificInputBase represents class SepaDirectDebitPaymentProduct771SpecificInputBase
type SepaDirectDebitPaymentProduct771SpecificInputBase struct {
	Mandate          *mandates.CreateMandateBase `json:"mandate,omitempty"`
	MandateReference *string                     `json:"mandateReference,omitempty"`
}

// NewSepaDirectDebitPaymentProduct771SpecificInputBase constructs a new SepaDirectDebitPaymentProduct771SpecificInputBase
func NewSepaDirectDebitPaymentProduct771SpecificInputBase() *SepaDirectDebitPaymentProduct771SpecificInputBase {
	return &SepaDirectDebitPaymentProduct771SpecificInputBase{}
}
