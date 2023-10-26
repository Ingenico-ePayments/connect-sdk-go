// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// PersonalInformation represents class PersonalInformation
type PersonalInformation struct {
	DateOfBirth    *string                 `json:"dateOfBirth,omitempty"`
	Gender         *string                 `json:"gender,omitempty"`
	Identification *PersonalIdentification `json:"identification,omitempty"`
	Name           *PersonalName           `json:"name,omitempty"`
}

// NewPersonalInformation constructs a new PersonalInformation
func NewPersonalInformation() *PersonalInformation {
	return &PersonalInformation{}
}
