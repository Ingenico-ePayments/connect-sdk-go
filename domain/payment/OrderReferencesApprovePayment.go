// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// OrderReferencesApprovePayment represents class OrderReferencesApprovePayment
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_OrderReferencesApprovePayment
type OrderReferencesApprovePayment struct {
	MerchantReference *string `json:"merchantReference,omitempty"`
}

// NewOrderReferencesApprovePayment constructs a new OrderReferencesApprovePayment
func NewOrderReferencesApprovePayment() *OrderReferencesApprovePayment {
	return &OrderReferencesApprovePayment{}
}
