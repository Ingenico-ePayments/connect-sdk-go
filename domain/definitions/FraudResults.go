// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package definitions

// FraudResults represents class FraudResults
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_FraudResults
type FraudResults struct {
	FraudServiceResult *string `json:"fraudServiceResult,omitempty"`
}

// NewFraudResults constructs a new FraudResults
func NewFraudResults() *FraudResults {
	return &FraudResults{}
}
