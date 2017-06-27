// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package riskassessments

// PersonalInformationRiskAssessment represents class PersonalInformationRiskAssessment
type PersonalInformationRiskAssessment struct {
	Name *PersonalNameRiskAssessment `json:"name,omitempty"`
}

// NewPersonalInformationRiskAssessment constructs a new PersonalInformationRiskAssessment
func NewPersonalInformationRiskAssessment() *PersonalInformationRiskAssessment {
	return &PersonalInformationRiskAssessment{}
}
