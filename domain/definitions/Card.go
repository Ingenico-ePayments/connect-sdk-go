// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package definitions

// Card represents class Card
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_Card
type Card struct {
	CardNumber     *string `json:"cardNumber,omitempty"`
	CardholderName *string `json:"cardholderName,omitempty"`
	Cvv            *string `json:"cvv,omitempty"`
	ExpiryDate     *string `json:"expiryDate,omitempty"`
	IssueNumber    *string `json:"issueNumber,omitempty"`
}

// NewCard constructs a new Card
func NewCard() *Card {
	return &Card{}
}
