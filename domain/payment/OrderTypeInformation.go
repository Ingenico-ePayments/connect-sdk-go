// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// OrderTypeInformation represents class OrderTypeInformation
type OrderTypeInformation struct {
	FundingType     *string `json:"fundingType,omitempty"`
	PurchaseType    *string `json:"purchaseType,omitempty"`
	TransactionType *string `json:"transactionType,omitempty"`
	UsageType       *string `json:"usageType,omitempty"`
}

// NewOrderTypeInformation constructs a new OrderTypeInformation
func NewOrderTypeInformation() *OrderTypeInformation {
	return &OrderTypeInformation{}
}
