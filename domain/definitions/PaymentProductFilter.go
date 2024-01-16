// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package definitions

// PaymentProductFilter represents class PaymentProductFilter
type PaymentProductFilter struct {
	Groups   *[]string `json:"groups,omitempty"`
	Products *[]int32  `json:"products,omitempty"`
}

// NewPaymentProductFilter constructs a new PaymentProductFilter
func NewPaymentProductFilter() *PaymentProductFilter {
	return &PaymentProductFilter{}
}
