// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package product

// PaymentProductGroups represents class PaymentProductGroups
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentProductGroups
type PaymentProductGroups struct {
	PaymentProductGroups *[]PaymentProductGroup `json:"paymentProductGroups,omitempty"`
}

// NewPaymentProductGroups constructs a new PaymentProductGroups
func NewPaymentProductGroups() *PaymentProductGroups {
	return &PaymentProductGroups{}
}
