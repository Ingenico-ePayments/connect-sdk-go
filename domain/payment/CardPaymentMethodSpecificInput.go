// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// CardPaymentMethodSpecificInput represents class CardPaymentMethodSpecificInput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_CardPaymentMethodSpecificInput
type CardPaymentMethodSpecificInput struct {
	AuthorizationMode                    *string                               `json:"authorizationMode,omitempty"`
	Card                                 *definitions.Card                     `json:"card,omitempty"`
	CustomerReference                    *string                               `json:"customerReference,omitempty"`
	ExternalCardholderAuthenticationData *ExternalCardholderAuthenticationData `json:"externalCardholderAuthenticationData,omitempty"`
	IsRecurring                          *bool                                 `json:"isRecurring,omitempty"`
	PaymentProductID                     *int32                                `json:"paymentProductId,omitempty"`
	RecurringPaymentSequenceIndicator    *string                               `json:"recurringPaymentSequenceIndicator,omitempty"`
	RequiresApproval                     *bool                                 `json:"requiresApproval,omitempty"`
	ReturnURL                            *string                               `json:"returnUrl,omitempty"`
	SkipAuthentication                   *bool                                 `json:"skipAuthentication,omitempty"`
	SkipFraudService                     *bool                                 `json:"skipFraudService,omitempty"`
	Token                                *string                               `json:"token,omitempty"`
	Tokenize                             *bool                                 `json:"tokenize,omitempty"`
}

// NewCardPaymentMethodSpecificInput constructs a new CardPaymentMethodSpecificInput
func NewCardPaymentMethodSpecificInput() *CardPaymentMethodSpecificInput {
	return &CardPaymentMethodSpecificInput{}
}
