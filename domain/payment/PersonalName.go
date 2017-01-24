// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// PersonalName represents class PersonalName
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PersonalName
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
