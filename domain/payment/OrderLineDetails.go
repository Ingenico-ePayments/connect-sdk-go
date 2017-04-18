// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// OrderLineDetails represents class OrderLineDetails
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_OrderLineDetails
type OrderLineDetails struct {
	DiscountAmount  *int64  `json:"discountAmount,omitempty"`
	LineAmountTotal *int64  `json:"lineAmountTotal,omitempty"`
	ProductCode     *string `json:"productCode,omitempty"`
	ProductPrice    *int64  `json:"productPrice,omitempty"`
	ProductType     *string `json:"productType,omitempty"`
	Quantity        *int64  `json:"quantity,omitempty"`
	TaxAmount       *int64  `json:"taxAmount,omitempty"`
	Unit            *string `json:"unit,omitempty"`
}

// NewOrderLineDetails constructs a new OrderLineDetails
func NewOrderLineDetails() *OrderLineDetails {
	return &OrderLineDetails{}
}
