// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// PersonalInformation represents class PersonalInformation
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PersonalInformation
type PersonalInformation struct {
	DateOfBirth *string       `json:"dateOfBirth,omitempty"`
	Gender      *string       `json:"gender,omitempty"`
	Name        *PersonalName `json:"name,omitempty"`
}

// NewPersonalInformation constructs a new PersonalInformation
func NewPersonalInformation() *PersonalInformation {
	return &PersonalInformation{}
}
