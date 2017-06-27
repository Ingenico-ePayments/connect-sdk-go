// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

// FraudResultsRetailDecisions represents class FraudResultsRetailDecisions
type FraudResultsRetailDecisions struct {
	FraudCode   *string `json:"fraudCode,omitempty"`
	FraudNeural *string `json:"fraudNeural,omitempty"`
	FraudRCF    *string `json:"fraudRCF,omitempty"`
}

// NewFraudResultsRetailDecisions constructs a new FraudResultsRetailDecisions
func NewFraudResultsRetailDecisions() *FraudResultsRetailDecisions {
	return &FraudResultsRetailDecisions{}
}
