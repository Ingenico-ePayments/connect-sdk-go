// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// OrderTypeInformation represents class OrderTypeInformation
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_OrderTypeInformation
type OrderTypeInformation struct {
	PurchaseType *string `json:"purchaseType,omitempty"`
	UsageType    *string `json:"usageType,omitempty"`
}

// NewOrderTypeInformation constructs a new OrderTypeInformation
func NewOrderTypeInformation() *OrderTypeInformation {
	return &OrderTypeInformation{}
}
