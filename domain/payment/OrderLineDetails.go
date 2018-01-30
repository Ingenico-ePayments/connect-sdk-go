// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// OrderLineDetails represents class OrderLineDetails
type OrderLineDetails struct {
	DiscountAmount          *int64  `json:"discountAmount,omitempty"`
	GoogleProductCategoryID *int64  `json:"googleProductCategoryId,omitempty"`
	LineAmountTotal         *int64  `json:"lineAmountTotal,omitempty"`
	ProductCategory         *string `json:"productCategory,omitempty"`
	ProductCode             *string `json:"productCode,omitempty"`
	ProductName             *string `json:"productName,omitempty"`
	ProductPrice            *int64  `json:"productPrice,omitempty"`
	ProductType             *string `json:"productType,omitempty"`
	Quantity                *int64  `json:"quantity,omitempty"`
	TaxAmount               *int64  `json:"taxAmount,omitempty"`
	Unit                    *string `json:"unit,omitempty"`
}

// NewOrderLineDetails constructs a new OrderLineDetails
func NewOrderLineDetails() *OrderLineDetails {
	return &OrderLineDetails{}
}
