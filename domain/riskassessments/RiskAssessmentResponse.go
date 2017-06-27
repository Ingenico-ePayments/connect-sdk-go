// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package riskassessments

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// RiskAssessmentResponse represents class RiskAssessmentResponse
type RiskAssessmentResponse struct {
	Results *[]definitions.ResultDoRiskAssessment `json:"results,omitempty"`
}

// NewRiskAssessmentResponse constructs a new RiskAssessmentResponse
func NewRiskAssessmentResponse() *RiskAssessmentResponse {
	return &RiskAssessmentResponse{}
}
