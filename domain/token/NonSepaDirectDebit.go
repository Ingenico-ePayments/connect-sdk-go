// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package token

// NonSepaDirectDebit represents class TokenNonSepaDirectDebit
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_TokenNonSepaDirectDebit
type NonSepaDirectDebit struct {
	Alias    *string                    `json:"alias,omitempty"`
	Customer *CustomerToken             `json:"customer,omitempty"`
	Mandate  *MandateNonSepaDirectDebit `json:"mandate,omitempty"`
}

// NewNonSepaDirectDebit constructs a new NonSepaDirectDebit
func NewNonSepaDirectDebit() *NonSepaDirectDebit {
	return &NonSepaDirectDebit{}
}
