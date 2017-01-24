// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package definitions

// PaymentProductFilter represents class PaymentProductFilter
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentProductFilter
type PaymentProductFilter struct {
	Groups   *[]string `json:"groups,omitempty"`
	Products *[]int32  `json:"products,omitempty"`
}

// NewPaymentProductFilter constructs a new PaymentProductFilter
func NewPaymentProductFilter() *PaymentProductFilter {
	return &PaymentProductFilter{}
}
