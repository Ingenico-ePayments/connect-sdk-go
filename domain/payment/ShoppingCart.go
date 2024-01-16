// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// ShoppingCart represents class ShoppingCart
type ShoppingCart struct {
	AmountBreakdown              *[]AmountBreakdown `json:"amountBreakdown,omitempty"`
	GiftCardPurchase             *GiftCardPurchase  `json:"giftCardPurchase,omitempty"`
	IsPreOrder                   *bool              `json:"isPreOrder,omitempty"`
	Items                        *[]LineItem        `json:"items,omitempty"`
	PreOrderItemAvailabilityDate *string            `json:"preOrderItemAvailabilityDate,omitempty"`
	ReOrderIndicator             *bool              `json:"reOrderIndicator,omitempty"`
}

// NewShoppingCart constructs a new ShoppingCart
func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{}
}
