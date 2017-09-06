// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// ShoppingCart represents class ShoppingCart
type ShoppingCart struct {
	AmountBreakdown *[]AmountBreakdown `json:"amountBreakdown,omitempty"`
	Items           *[]LineItem        `json:"items,omitempty"`
}

// NewShoppingCart constructs a new ShoppingCart
func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{}
}
