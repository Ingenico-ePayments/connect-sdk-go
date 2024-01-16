// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// SepaDirectDebitPaymentMethodSpecificOutput represents class SepaDirectDebitPaymentMethodSpecificOutput
type SepaDirectDebitPaymentMethodSpecificOutput struct {
	FraudResults                    *definitions.FraudResults `json:"fraudResults,omitempty"`
	PaymentProduct771SpecificOutput *Product771SpecificOutput `json:"paymentProduct771SpecificOutput,omitempty"`
	PaymentProductID                *int32                    `json:"paymentProductId,omitempty"`
	Token                           *string                   `json:"token,omitempty"`
}

// NewSepaDirectDebitPaymentMethodSpecificOutput constructs a new SepaDirectDebitPaymentMethodSpecificOutput
func NewSepaDirectDebitPaymentMethodSpecificOutput() *SepaDirectDebitPaymentMethodSpecificOutput {
	return &SepaDirectDebitPaymentMethodSpecificOutput{}
}
