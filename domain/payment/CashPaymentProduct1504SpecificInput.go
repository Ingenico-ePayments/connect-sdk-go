// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// CashPaymentProduct1504SpecificInput represents class CashPaymentProduct1504SpecificInput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_CashPaymentProduct1504SpecificInput
type CashPaymentProduct1504SpecificInput struct {
	ReturnURL *string `json:"returnUrl,omitempty"`
}

// NewCashPaymentProduct1504SpecificInput constructs a new CashPaymentProduct1504SpecificInput
func NewCashPaymentProduct1504SpecificInput() *CashPaymentProduct1504SpecificInput {
	return &CashPaymentProduct1504SpecificInput{}
}
