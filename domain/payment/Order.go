// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// Order represents class Order
type Order struct {
	AdditionalInput *AdditionalOrderInput      `json:"additionalInput,omitempty"`
	AmountOfMoney   *definitions.AmountOfMoney `json:"amountOfMoney,omitempty"`
	Customer        *Customer                  `json:"customer,omitempty"`
	// Deprecated: Use shoppingCart.items instead
	Items           *[]LineItem                `json:"items,omitempty"`
	References      *OrderReferences           `json:"references,omitempty"`
	// Deprecated: Use Merchant.seller instead
	Seller          *Seller                    `json:"seller,omitempty"`
	Shipping        *Shipping                  `json:"shipping,omitempty"`
	ShoppingCart    *ShoppingCart              `json:"shoppingCart,omitempty"`
}

// NewOrder constructs a new Order
func NewOrder() *Order {
	return &Order{}
}
