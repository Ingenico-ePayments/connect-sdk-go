// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package riskassessments

// PersonalNameRiskAssessment represents class PersonalNameRiskAssessment
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PersonalNameRiskAssessment
type PersonalNameRiskAssessment struct {
	FirstName     *string `json:"firstName,omitempty"`
	Surname       *string `json:"surname,omitempty"`
	SurnamePrefix *string `json:"surnamePrefix,omitempty"`
}

// NewPersonalNameRiskAssessment constructs a new PersonalNameRiskAssessment
func NewPersonalNameRiskAssessment() *PersonalNameRiskAssessment {
	return &PersonalNameRiskAssessment{}
}
