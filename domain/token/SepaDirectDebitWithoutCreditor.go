// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package token

// SepaDirectDebitWithoutCreditor represents class TokenSepaDirectDebitWithoutCreditor
type SepaDirectDebitWithoutCreditor struct {
	Alias    *string                                `json:"alias,omitempty"`
	Customer *CustomerTokenWithContactDetails       `json:"customer,omitempty"`
	Mandate  *MandateSepaDirectDebitWithoutCreditor `json:"mandate,omitempty"`
}

// NewSepaDirectDebitWithoutCreditor constructs a new SepaDirectDebitWithoutCreditor
func NewSepaDirectDebitWithoutCreditor() *SepaDirectDebitWithoutCreditor {
	return &SepaDirectDebitWithoutCreditor{}
}
