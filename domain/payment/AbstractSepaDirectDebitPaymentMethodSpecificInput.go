// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// AbstractSepaDirectDebitPaymentMethodSpecificInput represents class AbstractSepaDirectDebitPaymentMethodSpecificInput
type AbstractSepaDirectDebitPaymentMethodSpecificInput struct {
	PaymentProductID *int32 `json:"paymentProductId,omitempty"`
}

// NewAbstractSepaDirectDebitPaymentMethodSpecificInput constructs a new AbstractSepaDirectDebitPaymentMethodSpecificInput
func NewAbstractSepaDirectDebitPaymentMethodSpecificInput() *AbstractSepaDirectDebitPaymentMethodSpecificInput {
	return &AbstractSepaDirectDebitPaymentMethodSpecificInput{}
}
