// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package definitions

// CardFraudResults represents class CardFraudResults
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_CardFraudResults
type CardFraudResults struct {
	AvsResult          *string                      `json:"avsResult,omitempty"`
	CvvResult          *string                      `json:"cvvResult,omitempty"`
	FraudServiceResult *string                      `json:"fraudServiceResult,omitempty"`
	RetailDecisions    *FraudResultsRetailDecisions `json:"retailDecisions,omitempty"`
}

// NewCardFraudResults constructs a new CardFraudResults
func NewCardFraudResults() *CardFraudResults {
	return &CardFraudResults{}
}
