// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// CustomerPaymentActivity represents class CustomerPaymentActivity
type CustomerPaymentActivity struct {
	NumberOfPaymentAttemptsLast24Hours *int32 `json:"numberOfPaymentAttemptsLast24Hours,omitempty"`
	NumberOfPaymentAttemptsLastYear    *int32 `json:"numberOfPaymentAttemptsLastYear,omitempty"`
	NumberOfPurchasesLast6Months       *int32 `json:"numberOfPurchasesLast6Months,omitempty"`
}

// NewCustomerPaymentActivity constructs a new CustomerPaymentActivity
func NewCustomerPaymentActivity() *CustomerPaymentActivity {
	return &CustomerPaymentActivity{}
}
