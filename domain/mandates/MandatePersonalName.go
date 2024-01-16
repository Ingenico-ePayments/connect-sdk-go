// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package mandates

// MandatePersonalName represents class MandatePersonalName
type MandatePersonalName struct {
	FirstName *string `json:"firstName,omitempty"`
	Surname   *string `json:"surname,omitempty"`
}

// NewMandatePersonalName constructs a new MandatePersonalName
func NewMandatePersonalName() *MandatePersonalName {
	return &MandatePersonalName{}
}
