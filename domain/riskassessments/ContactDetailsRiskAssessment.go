// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package riskassessments

// ContactDetailsRiskAssessment represents class ContactDetailsRiskAssessment
type ContactDetailsRiskAssessment struct {
	EmailAddress *string `json:"emailAddress,omitempty"`
}

// NewContactDetailsRiskAssessment constructs a new ContactDetailsRiskAssessment
func NewContactDetailsRiskAssessment() *ContactDetailsRiskAssessment {
	return &ContactDetailsRiskAssessment{}
}
