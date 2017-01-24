// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package definitions

// ResultDoRiskAssessment represents class ResultDoRiskAssessment
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_ResultDoRiskAssessment
type ResultDoRiskAssessment struct {
	Category                          *string                            `json:"category,omitempty"`
	Result                            *string                            `json:"result,omitempty"`
	RetaildecisionsCCFraudCheckOutput *RetailDecisionsCCFraudCheckOutput `json:"retaildecisionsCCFraudCheckOutput,omitempty"`
	ValidationBankAccountOutput       *ValidationBankAccountOutput       `json:"validationBankAccountOutput,omitempty"`
}

// NewResultDoRiskAssessment constructs a new ResultDoRiskAssessment
func NewResultDoRiskAssessment() *ResultDoRiskAssessment {
	return &ResultDoRiskAssessment{}
}
