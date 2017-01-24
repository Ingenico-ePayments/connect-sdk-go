// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package riskassessments

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// RiskAssessmentResponse represents class RiskAssessmentResponse
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_RiskAssessmentResponse
type RiskAssessmentResponse struct {
	Results *[]definitions.ResultDoRiskAssessment `json:"results,omitempty"`
}

// NewRiskAssessmentResponse constructs a new RiskAssessmentResponse
func NewRiskAssessmentResponse() *RiskAssessmentResponse {
	return &RiskAssessmentResponse{}
}
