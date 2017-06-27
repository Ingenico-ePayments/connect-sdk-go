// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

// Address represents class Address
type Address struct {
	AdditionalInfo *string `json:"additionalInfo,omitempty"`
	City           *string `json:"city,omitempty"`
	CountryCode    *string `json:"countryCode,omitempty"`
	HouseNumber    *string `json:"houseNumber,omitempty"`
	State          *string `json:"state,omitempty"`
	StateCode      *string `json:"stateCode,omitempty"`
	Street         *string `json:"street,omitempty"`
	Zip            *string `json:"zip,omitempty"`
}

// NewAddress constructs a new Address
func NewAddress() *Address {
	return &Address{}
}
