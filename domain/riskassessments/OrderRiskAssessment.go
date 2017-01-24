// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package riskassessments

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// OrderRiskAssessment represents class OrderRiskAssessment
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_OrderRiskAssessment
type OrderRiskAssessment struct {
	AdditionalInput *definitions.AdditionalOrderInputAirlineData `json:"additionalInput,omitempty"`
	AmountOfMoney   *definitions.AmountOfMoney                   `json:"amountOfMoney,omitempty"`
	Customer        *CustomerRiskAssessment                      `json:"customer,omitempty"`
}

// NewOrderRiskAssessment constructs a new OrderRiskAssessment
func NewOrderRiskAssessment() *OrderRiskAssessment {
	return &OrderRiskAssessment{}
}
