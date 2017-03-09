// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// CardPaymentMethodSpecificInputBase represents class CardPaymentMethodSpecificInputBase
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_CardPaymentMethodSpecificInputBase
type CardPaymentMethodSpecificInputBase struct {
	AuthorizationMode                 *string `json:"authorizationMode,omitempty"`
	CustomerReference                 *string `json:"customerReference,omitempty"`
	PaymentProductID                  *int32  `json:"paymentProductId,omitempty"`
	RecurringPaymentSequenceIndicator *string `json:"recurringPaymentSequenceIndicator,omitempty"`
	RequiresApproval                  *bool   `json:"requiresApproval,omitempty"`
	SkipAuthentication                *bool   `json:"skipAuthentication,omitempty"`
	SkipFraudService                  *bool   `json:"skipFraudService,omitempty"`
	Token                             *string `json:"token,omitempty"`
	Tokenize                          *bool   `json:"tokenize,omitempty"`
}

// NewCardPaymentMethodSpecificInputBase constructs a new CardPaymentMethodSpecificInputBase
func NewCardPaymentMethodSpecificInputBase() *CardPaymentMethodSpecificInputBase {
	return &CardPaymentMethodSpecificInputBase{}
}
