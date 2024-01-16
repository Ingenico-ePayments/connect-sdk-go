// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// CashPaymentMethodSpecificInput represents class CashPaymentMethodSpecificInput
type CashPaymentMethodSpecificInput struct {
	// Deprecated: No replacement
	PaymentProduct1503SpecificInput *CashPaymentProduct1503SpecificInput `json:"paymentProduct1503SpecificInput,omitempty"`
	PaymentProduct1504SpecificInput *CashPaymentProduct1504SpecificInput `json:"paymentProduct1504SpecificInput,omitempty"`
	PaymentProduct1521SpecificInput *CashPaymentProduct1521SpecificInput `json:"paymentProduct1521SpecificInput,omitempty"`
	PaymentProduct1522SpecificInput *CashPaymentProduct1522SpecificInput `json:"paymentProduct1522SpecificInput,omitempty"`
	PaymentProduct1523SpecificInput *CashPaymentProduct1523SpecificInput `json:"paymentProduct1523SpecificInput,omitempty"`
	PaymentProduct1524SpecificInput *CashPaymentProduct1524SpecificInput `json:"paymentProduct1524SpecificInput,omitempty"`
	PaymentProduct1526SpecificInput *CashPaymentProduct1526SpecificInput `json:"paymentProduct1526SpecificInput,omitempty"`
	PaymentProductID                *int32                               `json:"paymentProductId,omitempty"`
}

// NewCashPaymentMethodSpecificInput constructs a new CashPaymentMethodSpecificInput
func NewCashPaymentMethodSpecificInput() *CashPaymentMethodSpecificInput {
	return &CashPaymentMethodSpecificInput{}
}
