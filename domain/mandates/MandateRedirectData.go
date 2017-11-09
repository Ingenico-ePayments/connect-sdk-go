// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package mandates

// MandateRedirectData represents class MandateRedirectData
type MandateRedirectData struct {
	RETURNMAC   *string `json:"RETURNMAC,omitempty"`
	RedirectURL *string `json:"redirectURL,omitempty"`
}

// NewMandateRedirectData constructs a new MandateRedirectData
func NewMandateRedirectData() *MandateRedirectData {
	return &MandateRedirectData{}
}
