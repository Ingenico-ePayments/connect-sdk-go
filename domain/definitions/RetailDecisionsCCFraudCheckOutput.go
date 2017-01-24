// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package definitions

// RetailDecisionsCCFraudCheckOutput represents class RetailDecisionsCCFraudCheckOutput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_RetailDecisionsCCFraudCheckOutput
type RetailDecisionsCCFraudCheckOutput struct {
	FraudCode   *string `json:"fraudCode,omitempty"`
	FraudNeural *string `json:"fraudNeural,omitempty"`
	FraudRCF    *string `json:"fraudRCF,omitempty"`
}

// NewRetailDecisionsCCFraudCheckOutput constructs a new RetailDecisionsCCFraudCheckOutput
func NewRetailDecisionsCCFraudCheckOutput() *RetailDecisionsCCFraudCheckOutput {
	return &RetailDecisionsCCFraudCheckOutput{}
}
