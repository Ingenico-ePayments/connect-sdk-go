// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// NonSepaDirectDebitPaymentMethodSpecificOutput represents class NonSepaDirectDebitPaymentMethodSpecificOutput
type NonSepaDirectDebitPaymentMethodSpecificOutput struct {
	FraudResults     *definitions.FraudResults `json:"fraudResults,omitempty"`
	PaymentProductID *int32                    `json:"paymentProductId,omitempty"`
}

// NewNonSepaDirectDebitPaymentMethodSpecificOutput constructs a new NonSepaDirectDebitPaymentMethodSpecificOutput
func NewNonSepaDirectDebitPaymentMethodSpecificOutput() *NonSepaDirectDebitPaymentMethodSpecificOutput {
	return &NonSepaDirectDebitPaymentMethodSpecificOutput{}
}
