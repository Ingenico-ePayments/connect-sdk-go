// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// CompletePaymentCardPaymentMethodSpecificInput represents class CompletePaymentCardPaymentMethodSpecificInput
type CompletePaymentCardPaymentMethodSpecificInput struct {
	Card *definitions.CardWithoutCvv `json:"card,omitempty"`
}

// NewCompletePaymentCardPaymentMethodSpecificInput constructs a new CompletePaymentCardPaymentMethodSpecificInput
func NewCompletePaymentCardPaymentMethodSpecificInput() *CompletePaymentCardPaymentMethodSpecificInput {
	return &CompletePaymentCardPaymentMethodSpecificInput{}
}
