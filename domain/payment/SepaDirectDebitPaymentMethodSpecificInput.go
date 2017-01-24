// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// SepaDirectDebitPaymentMethodSpecificInput represents class SepaDirectDebitPaymentMethodSpecificInput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_SepaDirectDebitPaymentMethodSpecificInput
type SepaDirectDebitPaymentMethodSpecificInput struct {
	DateCollect                       *string `json:"dateCollect,omitempty"`
	DirectDebitText                   *string `json:"directDebitText,omitempty"`
	IsRecurring                       *bool   `json:"isRecurring,omitempty"`
	PaymentProductID                  *int32  `json:"paymentProductId,omitempty"`
	RecurringPaymentSequenceIndicator *string `json:"recurringPaymentSequenceIndicator,omitempty"`
	Token                             *string `json:"token,omitempty"`
}

// NewSepaDirectDebitPaymentMethodSpecificInput constructs a new SepaDirectDebitPaymentMethodSpecificInput
func NewSepaDirectDebitPaymentMethodSpecificInput() *SepaDirectDebitPaymentMethodSpecificInput {
	return &SepaDirectDebitPaymentMethodSpecificInput{}
}
