// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// OrderOutput represents class OrderOutput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_OrderOutput
type OrderOutput struct {
	AmountOfMoney *definitions.AmountOfMoney `json:"amountOfMoney,omitempty"`
	References    *References                `json:"references,omitempty"`
}

// NewOrderOutput constructs a new OrderOutput
func NewOrderOutput() *OrderOutput {
	return &OrderOutput{}
}
