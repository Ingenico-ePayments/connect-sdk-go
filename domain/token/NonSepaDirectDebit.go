// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package token

// NonSepaDirectDebit represents class TokenNonSepaDirectDebit
type NonSepaDirectDebit struct {
	Alias    *string                    `json:"alias,omitempty"`
	Customer *CustomerToken             `json:"customer,omitempty"`
	Mandate  *MandateNonSepaDirectDebit `json:"mandate,omitempty"`
}

// NewNonSepaDirectDebit constructs a new NonSepaDirectDebit
func NewNonSepaDirectDebit() *NonSepaDirectDebit {
	return &NonSepaDirectDebit{}
}
