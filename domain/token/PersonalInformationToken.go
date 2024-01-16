// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package token

// PersonalInformationToken represents class PersonalInformationToken
type PersonalInformationToken struct {
	Name *PersonalNameToken `json:"name,omitempty"`
}

// NewPersonalInformationToken constructs a new PersonalInformationToken
func NewPersonalInformationToken() *PersonalInformationToken {
	return &PersonalInformationToken{}
}
