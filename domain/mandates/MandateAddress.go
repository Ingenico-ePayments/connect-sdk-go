// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package mandates

// MandateAddress represents class MandateAddress
type MandateAddress struct {
	City        *string `json:"city,omitempty"`
	CountryCode *string `json:"countryCode,omitempty"`
	HouseNumber *string `json:"houseNumber,omitempty"`
	Street      *string `json:"street,omitempty"`
	Zip         *string `json:"zip,omitempty"`
}

// NewMandateAddress constructs a new MandateAddress
func NewMandateAddress() *MandateAddress {
	return &MandateAddress{}
}
