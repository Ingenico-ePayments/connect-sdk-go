// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package riskassessments

// PersonalInformationRiskAssessment represents class PersonalInformationRiskAssessment
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PersonalInformationRiskAssessment
type PersonalInformationRiskAssessment struct {
	Name *PersonalNameRiskAssessment `json:"name,omitempty"`
}

// NewPersonalInformationRiskAssessment constructs a new PersonalInformationRiskAssessment
func NewPersonalInformationRiskAssessment() *PersonalInformationRiskAssessment {
	return &PersonalInformationRiskAssessment{}
}
