// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package mandates

// MandatePersonalInformation represents class MandatePersonalInformation
type MandatePersonalInformation struct {
	Name  *MandatePersonalName `json:"name,omitempty"`
	Title *string              `json:"title,omitempty"`
}

// NewMandatePersonalInformation constructs a new MandatePersonalInformation
func NewMandatePersonalInformation() *MandatePersonalInformation {
	return &MandatePersonalInformation{}
}
