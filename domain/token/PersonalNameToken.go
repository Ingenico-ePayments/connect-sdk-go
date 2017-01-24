// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package token

// PersonalNameToken represents class PersonalNameToken
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PersonalNameToken
type PersonalNameToken struct {
	FirstName     *string `json:"firstName,omitempty"`
	Surname       *string `json:"surname,omitempty"`
	SurnamePrefix *string `json:"surnamePrefix,omitempty"`
}

// NewPersonalNameToken constructs a new PersonalNameToken
func NewPersonalNameToken() *PersonalNameToken {
	return &PersonalNameToken{}
}
