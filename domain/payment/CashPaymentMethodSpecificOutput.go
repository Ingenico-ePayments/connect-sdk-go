// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// CashPaymentMethodSpecificOutput represents class CashPaymentMethodSpecificOutput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_CashPaymentMethodSpecificOutput
type CashPaymentMethodSpecificOutput struct {
	PaymentProductID *int32 `json:"paymentProductId,omitempty"`
}

// NewCashPaymentMethodSpecificOutput constructs a new CashPaymentMethodSpecificOutput
func NewCashPaymentMethodSpecificOutput() *CashPaymentMethodSpecificOutput {
	return &CashPaymentMethodSpecificOutput{}
}
