// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// CashPaymentMethodSpecificInputBase represents class CashPaymentMethodSpecificInputBase
type CashPaymentMethodSpecificInputBase struct {
	PaymentProductID *int32 `json:"paymentProductId,omitempty"`
}

// NewCashPaymentMethodSpecificInputBase constructs a new CashPaymentMethodSpecificInputBase
func NewCashPaymentMethodSpecificInputBase() *CashPaymentMethodSpecificInputBase {
	return &CashPaymentMethodSpecificInputBase{}
}
