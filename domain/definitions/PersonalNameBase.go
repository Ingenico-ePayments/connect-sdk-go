// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package definitions

// PersonalNameBase represents class PersonalNameBase
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PersonalNameBase
type PersonalNameBase struct {
	FirstName     *string `json:"firstName,omitempty"`
	Surname       *string `json:"surname,omitempty"`
	SurnamePrefix *string `json:"surnamePrefix,omitempty"`
}

// NewPersonalNameBase constructs a new PersonalNameBase
func NewPersonalNameBase() *PersonalNameBase {
	return &PersonalNameBase{}
}
