// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// SchemeTokenData represents class SchemeTokenData
type SchemeTokenData struct {
	CardholderName  *string `json:"cardholderName,omitempty"`
	Cryptogram      *string `json:"cryptogram,omitempty"`
	Eci             *string `json:"eci,omitempty"`
	NetworkToken    *string `json:"networkToken,omitempty"`
	TokenExpiryDate *string `json:"tokenExpiryDate,omitempty"`
}

// NewSchemeTokenData constructs a new SchemeTokenData
func NewSchemeTokenData() *SchemeTokenData {
	return &SchemeTokenData{}
}
