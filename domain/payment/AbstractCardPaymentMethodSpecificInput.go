// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// AbstractCardPaymentMethodSpecificInput represents class AbstractCardPaymentMethodSpecificInput
type AbstractCardPaymentMethodSpecificInput struct {
	AuthorizationMode                 *string `json:"authorizationMode,omitempty"`
	CustomerReference                 *string `json:"customerReference,omitempty"`
	PaymentProductID                  *int32  `json:"paymentProductId,omitempty"`
	RecurringPaymentSequenceIndicator *string `json:"recurringPaymentSequenceIndicator,omitempty"`
	RequiresApproval                  *bool   `json:"requiresApproval,omitempty"`
	SkipAuthentication                *bool   `json:"skipAuthentication,omitempty"`
	SkipFraudService                  *bool   `json:"skipFraudService,omitempty"`
	Token                             *string `json:"token,omitempty"`
	Tokenize                          *bool   `json:"tokenize,omitempty"`
	TransactionChannel                *string `json:"transactionChannel,omitempty"`
	UnscheduledCardOnFileIndicator    *string `json:"unscheduledCardOnFileIndicator,omitempty"`
	UnscheduledCardOnFileRequestor    *string `json:"unscheduledCardOnFileRequestor,omitempty"`
}

// NewAbstractCardPaymentMethodSpecificInput constructs a new AbstractCardPaymentMethodSpecificInput
func NewAbstractCardPaymentMethodSpecificInput() *AbstractCardPaymentMethodSpecificInput {
	return &AbstractCardPaymentMethodSpecificInput{}
}
