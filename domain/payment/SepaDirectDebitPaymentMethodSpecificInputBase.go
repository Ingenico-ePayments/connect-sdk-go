// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// SepaDirectDebitPaymentMethodSpecificInputBase represents class SepaDirectDebitPaymentMethodSpecificInputBase
type SepaDirectDebitPaymentMethodSpecificInputBase struct {
	PaymentProduct771SpecificInput *SepaDirectDebitPaymentProduct771SpecificInputBase `json:"paymentProduct771SpecificInput,omitempty"`
	PaymentProductID               *int32                                             `json:"paymentProductId,omitempty"`
}

// NewSepaDirectDebitPaymentMethodSpecificInputBase constructs a new SepaDirectDebitPaymentMethodSpecificInputBase
func NewSepaDirectDebitPaymentMethodSpecificInputBase() *SepaDirectDebitPaymentMethodSpecificInputBase {
	return &SepaDirectDebitPaymentMethodSpecificInputBase{}
}
