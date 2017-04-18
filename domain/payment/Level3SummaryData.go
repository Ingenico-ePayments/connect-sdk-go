// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// Level3SummaryData represents class Level3SummaryData
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_Level3SummaryData
//
// Deprecated: Use ShoppingCart instead
type Level3SummaryData struct {
	// Deprecated: Use ShoppingCart.discountAmount instead
	DiscountAmount *int64 `json:"discountAmount,omitempty"`
	// Deprecated: Use ShoppingCart.dutyAmount instead
	DutyAmount     *int64 `json:"dutyAmount,omitempty"`
	// Deprecated: Use ShoppingCart.shippingAmount instead
	ShippingAmount *int64 `json:"shippingAmount,omitempty"`
}

// NewLevel3SummaryData constructs a new Level3SummaryData
func NewLevel3SummaryData() *Level3SummaryData {
	return &Level3SummaryData{}
}
