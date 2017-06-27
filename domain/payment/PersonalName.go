// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// PersonalName represents class PersonalName
type PersonalName struct {
	FirstName     *string `json:"firstName,omitempty"`
	Surname       *string `json:"surname,omitempty"`
	SurnamePrefix *string `json:"surnamePrefix,omitempty"`
	Title         *string `json:"title,omitempty"`
}

// NewPersonalName constructs a new PersonalName
func NewPersonalName() *PersonalName {
	return &PersonalName{}
}
