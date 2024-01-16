// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// OrderOutput represents class OrderOutput
type OrderOutput struct {
	AmountOfMoney *definitions.AmountOfMoney `json:"amountOfMoney,omitempty"`
	References    *References                `json:"references,omitempty"`
}

// NewOrderOutput constructs a new OrderOutput
func NewOrderOutput() *OrderOutput {
	return &OrderOutput{}
}
