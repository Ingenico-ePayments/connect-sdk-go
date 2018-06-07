// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// CardPaymentMethodSpecificOutput represents class CardPaymentMethodSpecificOutput
type CardPaymentMethodSpecificOutput struct {
	AuthorisationCode   *string                       `json:"authorisationCode,omitempty"`
	Card                *definitions.CardEssentials   `json:"card,omitempty"`
	FraudResults        *definitions.CardFraudResults `json:"fraudResults,omitempty"`
	PaymentProductID    *int32                        `json:"paymentProductId,omitempty"`
	ThreeDSecureResults *ThreeDSecureResults          `json:"threeDSecureResults,omitempty"`
	Token               *string                       `json:"token,omitempty"`
}

// NewCardPaymentMethodSpecificOutput constructs a new CardPaymentMethodSpecificOutput
func NewCardPaymentMethodSpecificOutput() *CardPaymentMethodSpecificOutput {
	return &CardPaymentMethodSpecificOutput{}
}
