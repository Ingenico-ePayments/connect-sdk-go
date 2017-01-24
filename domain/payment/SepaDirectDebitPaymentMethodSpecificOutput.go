// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// SepaDirectDebitPaymentMethodSpecificOutput represents class SepaDirectDebitPaymentMethodSpecificOutput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_SepaDirectDebitPaymentMethodSpecificOutput
type SepaDirectDebitPaymentMethodSpecificOutput struct {
	FraudResults     *definitions.FraudResults `json:"fraudResults,omitempty"`
	PaymentProductID *int32                    `json:"paymentProductId,omitempty"`
}

// NewSepaDirectDebitPaymentMethodSpecificOutput constructs a new SepaDirectDebitPaymentMethodSpecificOutput
func NewSepaDirectDebitPaymentMethodSpecificOutput() *SepaDirectDebitPaymentMethodSpecificOutput {
	return &SepaDirectDebitPaymentMethodSpecificOutput{}
}
