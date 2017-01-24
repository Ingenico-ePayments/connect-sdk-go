// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package token

// Debtor represents class Debtor
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_Debtor
type Debtor struct {
	AdditionalAddressInfo *string `json:"additionalAddressInfo,omitempty"`
	City                  *string `json:"city,omitempty"`
	CountryCode           *string `json:"countryCode,omitempty"`
	FirstName             *string `json:"firstName,omitempty"`
	HouseNumber           *string `json:"houseNumber,omitempty"`
	State                 *string `json:"state,omitempty"`
	StateCode             *string `json:"stateCode,omitempty"`
	Street                *string `json:"street,omitempty"`
	Surname               *string `json:"surname,omitempty"`
	SurnamePrefix         *string `json:"surnamePrefix,omitempty"`
	Zip                   *string `json:"zip,omitempty"`
}

// NewDebtor constructs a new Debtor
func NewDebtor() *Debtor {
	return &Debtor{}
}
