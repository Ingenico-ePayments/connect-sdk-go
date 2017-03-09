// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// NonSepaDirectDebitPaymentMethodSpecificInput represents class NonSepaDirectDebitPaymentMethodSpecificInput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_NonSepaDirectDebitPaymentMethodSpecificInput
type NonSepaDirectDebitPaymentMethodSpecificInput struct {
	DateCollect                       *string                                           `json:"dateCollect,omitempty"`
	DirectDebitText                   *string                                           `json:"directDebitText,omitempty"`
	IsRecurring                       *bool                                             `json:"isRecurring,omitempty"`
	PaymentProduct705SpecificInput    *NonSepaDirectDebitPaymentProduct705SpecificInput `json:"paymentProduct705SpecificInput,omitempty"`
	PaymentProductID                  *int32                                            `json:"paymentProductId,omitempty"`
	RecurringPaymentSequenceIndicator *string                                           `json:"recurringPaymentSequenceIndicator,omitempty"`
	Token                             *string                                           `json:"token,omitempty"`
	Tokenize                          *bool                                             `json:"tokenize,omitempty"`
}

// NewNonSepaDirectDebitPaymentMethodSpecificInput constructs a new NonSepaDirectDebitPaymentMethodSpecificInput
func NewNonSepaDirectDebitPaymentMethodSpecificInput() *NonSepaDirectDebitPaymentMethodSpecificInput {
	return &NonSepaDirectDebitPaymentMethodSpecificInput{}
}
