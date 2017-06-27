// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package riskassessments

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// OrderRiskAssessment represents class OrderRiskAssessment
type OrderRiskAssessment struct {
	AdditionalInput *definitions.AdditionalOrderInputAirlineData `json:"additionalInput,omitempty"`
	AmountOfMoney   *definitions.AmountOfMoney                   `json:"amountOfMoney,omitempty"`
	Customer        *CustomerRiskAssessment                      `json:"customer,omitempty"`
}

// NewOrderRiskAssessment constructs a new OrderRiskAssessment
func NewOrderRiskAssessment() *OrderRiskAssessment {
	return &OrderRiskAssessment{}
}
