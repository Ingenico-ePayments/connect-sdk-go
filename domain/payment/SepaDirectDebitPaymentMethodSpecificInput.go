// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// SepaDirectDebitPaymentMethodSpecificInput represents class SepaDirectDebitPaymentMethodSpecificInput
type SepaDirectDebitPaymentMethodSpecificInput struct {
	DateCollect                       *string                                        `json:"dateCollect,omitempty"`
	DirectDebitText                   *string                                        `json:"directDebitText,omitempty"`
	IsRecurring                       *bool                                          `json:"isRecurring,omitempty"`
	PaymentProduct771SpecificInput    *SepaDirectDebitPaymentProduct771SpecificInput `json:"paymentProduct771SpecificInput,omitempty"`
	PaymentProductID                  *int32                                         `json:"paymentProductId,omitempty"`
	RecurringPaymentSequenceIndicator *string                                        `json:"recurringPaymentSequenceIndicator,omitempty"`
	Token                             *string                                        `json:"token,omitempty"`
	Tokenize                          *bool                                          `json:"tokenize,omitempty"`
}

// NewSepaDirectDebitPaymentMethodSpecificInput constructs a new SepaDirectDebitPaymentMethodSpecificInput
func NewSepaDirectDebitPaymentMethodSpecificInput() *SepaDirectDebitPaymentMethodSpecificInput {
	return &SepaDirectDebitPaymentMethodSpecificInput{}
}
