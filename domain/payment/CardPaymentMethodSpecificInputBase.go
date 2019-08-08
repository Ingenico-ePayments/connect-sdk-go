// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// CardPaymentMethodSpecificInputBase represents class CardPaymentMethodSpecificInputBase
type CardPaymentMethodSpecificInputBase struct {
	AuthorizationMode                      *string                `json:"authorizationMode,omitempty"`
	CustomerReference                      *string                `json:"customerReference,omitempty"`
	InitialSchemeTransactionID             *string                `json:"initialSchemeTransactionId,omitempty"`
	PaymentProductID                       *int32                 `json:"paymentProductId,omitempty"`
	Recurring                              *CardRecurrenceDetails `json:"recurring,omitempty"`
	// Deprecated: Use recurring.recurringPaymentSequenceIndicator instead
	RecurringPaymentSequenceIndicator      *string                `json:"recurringPaymentSequenceIndicator,omitempty"`
	RequiresApproval                       *bool                  `json:"requiresApproval,omitempty"`
	// Deprecated: Use threeDSecure.skipAuthentication instead
	SkipAuthentication                     *bool                  `json:"skipAuthentication,omitempty"`
	SkipFraudService                       *bool                  `json:"skipFraudService,omitempty"`
	ThreeDSecure                           *ThreeDSecureBase      `json:"threeDSecure,omitempty"`
	Token                                  *string                `json:"token,omitempty"`
	Tokenize                               *bool                  `json:"tokenize,omitempty"`
	TransactionChannel                     *string                `json:"transactionChannel,omitempty"`
	// Deprecated: Use unscheduledCardOnFileSequenceIndicator instead
	UnscheduledCardOnFileIndicator         *string                `json:"unscheduledCardOnFileIndicator,omitempty"`
	UnscheduledCardOnFileRequestor         *string                `json:"unscheduledCardOnFileRequestor,omitempty"`
	UnscheduledCardOnFileSequenceIndicator *string                `json:"unscheduledCardOnFileSequenceIndicator,omitempty"`
}

// NewCardPaymentMethodSpecificInputBase constructs a new CardPaymentMethodSpecificInputBase
func NewCardPaymentMethodSpecificInputBase() *CardPaymentMethodSpecificInputBase {
	return &CardPaymentMethodSpecificInputBase{}
}
