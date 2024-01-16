// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// OrderReferencesApprovePayment represents class OrderReferencesApprovePayment
type OrderReferencesApprovePayment struct {
	MerchantReference *string `json:"merchantReference,omitempty"`
}

// NewOrderReferencesApprovePayment constructs a new OrderReferencesApprovePayment
func NewOrderReferencesApprovePayment() *OrderReferencesApprovePayment {
	return &OrderReferencesApprovePayment{}
}
