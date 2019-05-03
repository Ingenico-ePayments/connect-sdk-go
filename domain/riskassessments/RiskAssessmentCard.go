// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package riskassessments

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// RiskAssessmentCard represents class RiskAssessmentCard
type RiskAssessmentCard struct {
	Card             *definitions.Card        `json:"card,omitempty"`
	FraudFields      *definitions.FraudFields `json:"fraudFields,omitempty"`
	Merchant         *MerchantRiskAssessment  `json:"merchant,omitempty"`
	Order            *OrderRiskAssessment     `json:"order,omitempty"`
	PaymentProductID *int32                   `json:"paymentProductId,omitempty"`
}

// NewRiskAssessmentCard constructs a new RiskAssessmentCard
func NewRiskAssessmentCard() *RiskAssessmentCard {
	return &RiskAssessmentCard{}
}
