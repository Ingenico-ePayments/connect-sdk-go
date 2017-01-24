// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package definitions

// FraudResultsRetailDecisions represents class FraudResultsRetailDecisions
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_FraudResultsRetailDecisions
type FraudResultsRetailDecisions struct {
	FraudCode   *string `json:"fraudCode,omitempty"`
	FraudNeural *string `json:"fraudNeural,omitempty"`
	FraudRCF    *string `json:"fraudRCF,omitempty"`
}

// NewFraudResultsRetailDecisions constructs a new FraudResultsRetailDecisions
func NewFraudResultsRetailDecisions() *FraudResultsRetailDecisions {
	return &FraudResultsRetailDecisions{}
}
