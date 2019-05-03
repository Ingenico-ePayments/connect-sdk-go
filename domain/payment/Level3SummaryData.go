// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// Level3SummaryData represents class Level3SummaryData
//
// Deprecated: Use ShoppingCart.amountBreakdown instead
type Level3SummaryData struct {
	// Deprecated: Use ShoppingCart.amountBreakdown with type DISCOUNT instead
	DiscountAmount *int64 `json:"discountAmount,omitempty"`
	// Deprecated: Use ShoppingCart.amountBreakdown with type DUTY instead
	DutyAmount     *int64 `json:"dutyAmount,omitempty"`
	// Deprecated: Use ShoppingCart.amountBreakdown with type SHIPPING instead
	ShippingAmount *int64 `json:"shippingAmount,omitempty"`
}

// NewLevel3SummaryData constructs a new Level3SummaryData
func NewLevel3SummaryData() *Level3SummaryData {
	return &Level3SummaryData{}
}
