// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

// RetailDecisionsCCFraudCheckOutput represents class RetailDecisionsCCFraudCheckOutput
type RetailDecisionsCCFraudCheckOutput struct {
	FraudCode   *string `json:"fraudCode,omitempty"`
	FraudNeural *string `json:"fraudNeural,omitempty"`
	FraudRCF    *string `json:"fraudRCF,omitempty"`
}

// NewRetailDecisionsCCFraudCheckOutput constructs a new RetailDecisionsCCFraudCheckOutput
func NewRetailDecisionsCCFraudCheckOutput() *RetailDecisionsCCFraudCheckOutput {
	return &RetailDecisionsCCFraudCheckOutput{}
}
