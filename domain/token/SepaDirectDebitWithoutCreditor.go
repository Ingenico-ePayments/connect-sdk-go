// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package token

// SepaDirectDebitWithoutCreditor represents class TokenSepaDirectDebitWithoutCreditor
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_TokenSepaDirectDebitWithoutCreditor
type SepaDirectDebitWithoutCreditor struct {
	Alias    *string                                `json:"alias,omitempty"`
	Customer *CustomerTokenWithContactDetails       `json:"customer,omitempty"`
	Mandate  *MandateSepaDirectDebitWithoutCreditor `json:"mandate,omitempty"`
}

// NewSepaDirectDebitWithoutCreditor constructs a new SepaDirectDebitWithoutCreditor
func NewSepaDirectDebitWithoutCreditor() *SepaDirectDebitWithoutCreditor {
	return &SepaDirectDebitWithoutCreditor{}
}
