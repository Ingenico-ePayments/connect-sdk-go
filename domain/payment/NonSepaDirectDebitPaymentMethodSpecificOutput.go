// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// NonSepaDirectDebitPaymentMethodSpecificOutput represents class NonSepaDirectDebitPaymentMethodSpecificOutput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_NonSepaDirectDebitPaymentMethodSpecificOutput
type NonSepaDirectDebitPaymentMethodSpecificOutput struct {
	FraudResults     *definitions.FraudResults `json:"fraudResults,omitempty"`
	PaymentProductID *int32                    `json:"paymentProductId,omitempty"`
}

// NewNonSepaDirectDebitPaymentMethodSpecificOutput constructs a new NonSepaDirectDebitPaymentMethodSpecificOutput
func NewNonSepaDirectDebitPaymentMethodSpecificOutput() *NonSepaDirectDebitPaymentMethodSpecificOutput {
	return &NonSepaDirectDebitPaymentMethodSpecificOutput{}
}
