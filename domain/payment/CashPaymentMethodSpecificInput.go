// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// CashPaymentMethodSpecificInput represents class CashPaymentMethodSpecificInput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_CashPaymentMethodSpecificInput
type CashPaymentMethodSpecificInput struct {
	PaymentProduct1503SpecificInput *CashPaymentProduct1503SpecificInput `json:"paymentProduct1503SpecificInput,omitempty"`
	PaymentProduct1504SpecificInput *CashPaymentProduct1504SpecificInput `json:"paymentProduct1504SpecificInput,omitempty"`
	PaymentProductID                *int32                               `json:"paymentProductId,omitempty"`
}

// NewCashPaymentMethodSpecificInput constructs a new CashPaymentMethodSpecificInput
func NewCashPaymentMethodSpecificInput() *CashPaymentMethodSpecificInput {
	return &CashPaymentMethodSpecificInput{}
}
