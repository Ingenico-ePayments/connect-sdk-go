// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package mandates

// MandateContactDetails represents class MandateContactDetails
type MandateContactDetails struct {
	EmailAddress *string `json:"emailAddress,omitempty"`
}

// NewMandateContactDetails constructs a new MandateContactDetails
func NewMandateContactDetails() *MandateContactDetails {
	return &MandateContactDetails{}
}
