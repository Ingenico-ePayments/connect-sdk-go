// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// CashPaymentMethodSpecificInputBase represents class CashPaymentMethodSpecificInputBase
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_CashPaymentMethodSpecificInputBase
type CashPaymentMethodSpecificInputBase struct {
	PaymentProductID *int32 `json:"paymentProductId,omitempty"`
}

// NewCashPaymentMethodSpecificInputBase constructs a new CashPaymentMethodSpecificInputBase
func NewCashPaymentMethodSpecificInputBase() *CashPaymentMethodSpecificInputBase {
	return &CashPaymentMethodSpecificInputBase{}
}
