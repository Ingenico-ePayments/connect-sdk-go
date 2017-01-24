// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// CashPaymentProduct1503SpecificInput represents class CashPaymentProduct1503SpecificInput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_CashPaymentProduct1503SpecificInput
type CashPaymentProduct1503SpecificInput struct {
	ReturnURL *string `json:"returnUrl,omitempty"`
}

// NewCashPaymentProduct1503SpecificInput constructs a new CashPaymentProduct1503SpecificInput
func NewCashPaymentProduct1503SpecificInput() *CashPaymentProduct1503SpecificInput {
	return &CashPaymentProduct1503SpecificInput{}
}
