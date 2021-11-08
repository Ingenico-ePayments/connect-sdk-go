// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

// MicrosoftFraudResults represents class MicrosoftFraudResults
type MicrosoftFraudResults struct {
	FraudScore *int32 `json:"fraudScore,omitempty"`
}

// NewMicrosoftFraudResults constructs a new MicrosoftFraudResults
func NewMicrosoftFraudResults() *MicrosoftFraudResults {
	return &MicrosoftFraudResults{}
}
