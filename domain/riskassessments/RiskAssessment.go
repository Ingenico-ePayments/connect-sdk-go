// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package riskassessments

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// RiskAssessment represents class RiskAssessment
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_RiskAssessment
type RiskAssessment struct {
	FraudFields      *definitions.FraudFields `json:"fraudFields,omitempty"`
	Order            *OrderRiskAssessment     `json:"order,omitempty"`
	PaymentProductID *int32                   `json:"paymentProductId,omitempty"`
}

// NewRiskAssessment constructs a new RiskAssessment
func NewRiskAssessment() *RiskAssessment {
	return &RiskAssessment{}
}
