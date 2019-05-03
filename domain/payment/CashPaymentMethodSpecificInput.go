// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// CashPaymentMethodSpecificInput represents class CashPaymentMethodSpecificInput
type CashPaymentMethodSpecificInput struct {
	// Deprecated: No replacement
	PaymentProduct1503SpecificInput *CashPaymentProduct1503SpecificInput `json:"paymentProduct1503SpecificInput,omitempty"`
	PaymentProduct1504SpecificInput *CashPaymentProduct1504SpecificInput `json:"paymentProduct1504SpecificInput,omitempty"`
	PaymentProductID                *int32                               `json:"paymentProductId,omitempty"`
}

// NewCashPaymentMethodSpecificInput constructs a new CashPaymentMethodSpecificInput
func NewCashPaymentMethodSpecificInput() *CashPaymentMethodSpecificInput {
	return &CashPaymentMethodSpecificInput{}
}
