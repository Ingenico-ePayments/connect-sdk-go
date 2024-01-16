// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// CardPaymentMethodSpecificInput represents class CardPaymentMethodSpecificInput
type CardPaymentMethodSpecificInput struct {
	AcquirerPromotionCode                  *string                               `json:"acquirerPromotionCode,omitempty"`
	AuthorizationMode                      *string                               `json:"authorizationMode,omitempty"`
	Card                                   *definitions.Card                     `json:"card,omitempty"`
	CustomerReference                      *string                               `json:"customerReference,omitempty"`
	// Deprecated: Use threeDSecure.externalCardholderAuthenticationData instead
	ExternalCardholderAuthenticationData   *ExternalCardholderAuthenticationData `json:"externalCardholderAuthenticationData,omitempty"`
	InitialSchemeTransactionID             *string                               `json:"initialSchemeTransactionId,omitempty"`
	IsRecurring                            *bool                                 `json:"isRecurring,omitempty"`
	MerchantInitiatedReasonIndicator       *string                               `json:"merchantInitiatedReasonIndicator,omitempty"`
	NetworkTokenData                       *SchemeTokenData                      `json:"networkTokenData,omitempty"`
	PaymentProductID                       *int32                                `json:"paymentProductId,omitempty"`
	Recurring                              *CardRecurrenceDetails                `json:"recurring,omitempty"`
	// Deprecated: Use recurring.recurringPaymentSequenceIndicator instead
	RecurringPaymentSequenceIndicator      *string                               `json:"recurringPaymentSequenceIndicator,omitempty"`
	RequiresApproval                       *bool                                 `json:"requiresApproval,omitempty"`
	// Deprecated: Use threeDSecure.redirectionData.returnUrl instead
	ReturnURL                              *string                               `json:"returnUrl,omitempty"`
	// Deprecated: Use threeDSecure.skipAuthentication instead
	SkipAuthentication                     *bool                                 `json:"skipAuthentication,omitempty"`
	SkipFraudService                       *bool                                 `json:"skipFraudService,omitempty"`
	ThreeDSecure                           *ThreeDSecure                         `json:"threeDSecure,omitempty"`
	Token                                  *string                               `json:"token,omitempty"`
	Tokenize                               *bool                                 `json:"tokenize,omitempty"`
	TransactionChannel                     *string                               `json:"transactionChannel,omitempty"`
	// Deprecated: Use unscheduledCardOnFileSequenceIndicator instead
	UnscheduledCardOnFileIndicator         *string                               `json:"unscheduledCardOnFileIndicator,omitempty"`
	UnscheduledCardOnFileRequestor         *string                               `json:"unscheduledCardOnFileRequestor,omitempty"`
	UnscheduledCardOnFileSequenceIndicator *string                               `json:"unscheduledCardOnFileSequenceIndicator,omitempty"`
}

// NewCardPaymentMethodSpecificInput constructs a new CardPaymentMethodSpecificInput
func NewCardPaymentMethodSpecificInput() *CardPaymentMethodSpecificInput {
	return &CardPaymentMethodSpecificInput{}
}
