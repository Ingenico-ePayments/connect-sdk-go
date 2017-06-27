// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package token

// PersonalNameToken represents class PersonalNameToken
type PersonalNameToken struct {
	FirstName     *string `json:"firstName,omitempty"`
	Surname       *string `json:"surname,omitempty"`
	SurnamePrefix *string `json:"surnamePrefix,omitempty"`
}

// NewPersonalNameToken constructs a new PersonalNameToken
func NewPersonalNameToken() *PersonalNameToken {
	return &PersonalNameToken{}
}
