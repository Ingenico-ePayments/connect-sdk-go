// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package riskassessments

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// RiskAssessmentCard represents class RiskAssessmentCard
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_RiskAssessmentCard
type RiskAssessmentCard struct {
	Card             *definitions.Card        `json:"card,omitempty"`
	FraudFields      *definitions.FraudFields `json:"fraudFields,omitempty"`
	Order            *OrderRiskAssessment     `json:"order,omitempty"`
	PaymentProductID *int32                   `json:"paymentProductId,omitempty"`
}

// NewRiskAssessmentCard constructs a new RiskAssessmentCard
func NewRiskAssessmentCard() *RiskAssessmentCard {
	return &RiskAssessmentCard{}
}
