// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// OrderApprovePayment represents class OrderApprovePayment
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_OrderApprovePayment
type OrderApprovePayment struct {
	AdditionalInput *definitions.AdditionalOrderInputAirlineData `json:"additionalInput,omitempty"`
	References      *OrderReferencesApprovePayment               `json:"references,omitempty"`
}

// NewOrderApprovePayment constructs a new OrderApprovePayment
func NewOrderApprovePayment() *OrderApprovePayment {
	return &OrderApprovePayment{}
}
