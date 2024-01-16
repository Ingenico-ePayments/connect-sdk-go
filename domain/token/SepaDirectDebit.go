// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package token

// SepaDirectDebit represents class TokenSepaDirectDebit
type SepaDirectDebit struct {
	Alias    *string                          `json:"alias,omitempty"`
	Customer *CustomerTokenWithContactDetails `json:"customer,omitempty"`
	Mandate  *MandateSepaDirectDebit          `json:"mandate,omitempty"`
}

// NewSepaDirectDebit constructs a new SepaDirectDebit
func NewSepaDirectDebit() *SepaDirectDebit {
	return &SepaDirectDebit{}
}
