// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// OrderApprovePayment represents class OrderApprovePayment
type OrderApprovePayment struct {
	AdditionalInput *definitions.AdditionalOrderInputAirlineData `json:"additionalInput,omitempty"`
	Customer        *CustomerApprovePayment                      `json:"customer,omitempty"`
	References      *OrderReferencesApprovePayment               `json:"references,omitempty"`
}

// NewOrderApprovePayment constructs a new OrderApprovePayment
func NewOrderApprovePayment() *OrderApprovePayment {
	return &OrderApprovePayment{}
}
