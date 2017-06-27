// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// OrderTypeInformation represents class OrderTypeInformation
type OrderTypeInformation struct {
	PurchaseType *string `json:"purchaseType,omitempty"`
	UsageType    *string `json:"usageType,omitempty"`
}

// NewOrderTypeInformation constructs a new OrderTypeInformation
func NewOrderTypeInformation() *OrderTypeInformation {
	return &OrderTypeInformation{}
}
