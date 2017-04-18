// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// ShoppingCart represents class ShoppingCart
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_ShoppingCart
type ShoppingCart struct {
	AmountBreakdown *[]AmountBreakdown `json:"amountBreakdown,omitempty"`
}

// NewShoppingCart constructs a new ShoppingCart
func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{}
}
