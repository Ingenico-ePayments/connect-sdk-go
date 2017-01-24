// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package token

// SepaDirectDebit represents class TokenSepaDirectDebit
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_TokenSepaDirectDebit
type SepaDirectDebit struct {
	Alias    *string                          `json:"alias,omitempty"`
	Customer *CustomerTokenWithContactDetails `json:"customer,omitempty"`
	Mandate  *MandateSepaDirectDebit          `json:"mandate,omitempty"`
}

// NewSepaDirectDebit constructs a new SepaDirectDebit
func NewSepaDirectDebit() *SepaDirectDebit {
	return &SepaDirectDebit{}
}
