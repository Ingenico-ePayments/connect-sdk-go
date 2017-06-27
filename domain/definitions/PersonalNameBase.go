// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

// PersonalNameBase represents class PersonalNameBase
type PersonalNameBase struct {
	FirstName     *string `json:"firstName,omitempty"`
	Surname       *string `json:"surname,omitempty"`
	SurnamePrefix *string `json:"surnamePrefix,omitempty"`
}

// NewPersonalNameBase constructs a new PersonalNameBase
func NewPersonalNameBase() *PersonalNameBase {
	return &PersonalNameBase{}
}
