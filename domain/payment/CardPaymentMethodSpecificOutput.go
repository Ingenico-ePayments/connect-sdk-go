// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// CardPaymentMethodSpecificOutput represents class CardPaymentMethodSpecificOutput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_CardPaymentMethodSpecificOutput
type CardPaymentMethodSpecificOutput struct {
	AuthorisationCode   *string                       `json:"authorisationCode,omitempty"`
	Card                *definitions.CardEssentials   `json:"card,omitempty"`
	FraudResults        *definitions.CardFraudResults `json:"fraudResults,omitempty"`
	PaymentProductID    *int32                        `json:"paymentProductId,omitempty"`
	ThreeDSecureResults *ThreeDSecureResults          `json:"threeDSecureResults,omitempty"`
}

// NewCardPaymentMethodSpecificOutput constructs a new CardPaymentMethodSpecificOutput
func NewCardPaymentMethodSpecificOutput() *CardPaymentMethodSpecificOutput {
	return &CardPaymentMethodSpecificOutput{}
}
