// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

// FraugsterResults represents class FraugsterResults
type FraugsterResults struct {
	FraudInvestigationPoints *string `json:"fraudInvestigationPoints,omitempty"`
	FraudScore               *int32  `json:"fraudScore,omitempty"`
}

// NewFraugsterResults constructs a new FraugsterResults
func NewFraugsterResults() *FraugsterResults {
	return &FraugsterResults{}
}
