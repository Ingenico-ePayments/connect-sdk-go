// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// Order represents class Order
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_Order
type Order struct {
	AdditionalInput *AdditionalOrderInput      `json:"additionalInput,omitempty"`
	AmountOfMoney   *definitions.AmountOfMoney `json:"amountOfMoney,omitempty"`
	Customer        *Customer                  `json:"customer,omitempty"`
	Items           *[]LineItem                `json:"items,omitempty"`
	References      *OrderReferences           `json:"references,omitempty"`
	ShoppingCart    *ShoppingCart              `json:"shoppingCart,omitempty"`
}

// NewOrder constructs a new Order
func NewOrder() *Order {
	return &Order{}
}
