// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

// FraudResults represents class FraudResults
type FraudResults struct {
	FraudServiceResult *string `json:"fraudServiceResult,omitempty"`
	InAuth             *InAuth `json:"inAuth,omitempty"`
}

// NewFraudResults constructs a new FraudResults
func NewFraudResults() *FraudResults {
	return &FraudResults{}
}
