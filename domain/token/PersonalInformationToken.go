// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package token

// PersonalInformationToken represents class PersonalInformationToken
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PersonalInformationToken
type PersonalInformationToken struct {
	Name *PersonalNameToken `json:"name,omitempty"`
}

// NewPersonalInformationToken constructs a new PersonalInformationToken
func NewPersonalInformationToken() *PersonalInformationToken {
	return &PersonalInformationToken{}
}
