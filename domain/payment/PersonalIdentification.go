// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// PersonalIdentification represents class PersonalIdentification
type PersonalIdentification struct {
	IDIssuingCountryCode *string `json:"idIssuingCountryCode,omitempty"`
	IDType               *string `json:"idType,omitempty"`
	IDValue              *string `json:"idValue,omitempty"`
}

// NewPersonalIdentification constructs a new PersonalIdentification
func NewPersonalIdentification() *PersonalIdentification {
	return &PersonalIdentification{}
}
