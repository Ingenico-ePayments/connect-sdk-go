// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// AdditionalOrderInput represents class AdditionalOrderInput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_AdditionalOrderInput
type AdditionalOrderInput struct {
	AirlineData          *definitions.AirlineData `json:"airlineData,omitempty"`
	// Deprecated: Use Order.shoppingCart instead
	Level3SummaryData    *Level3SummaryData       `json:"level3SummaryData,omitempty"`
	NumberOfInstallments *int64                   `json:"numberOfInstallments,omitempty"`
	OrderDate            *string                  `json:"orderDate,omitempty"`
	TypeInformation      *OrderTypeInformation    `json:"typeInformation,omitempty"`
}

// NewAdditionalOrderInput constructs a new AdditionalOrderInput
func NewAdditionalOrderInput() *AdditionalOrderInput {
	return &AdditionalOrderInput{}
}
