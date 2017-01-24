// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package token

// Creditor represents class Creditor
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_Creditor
type Creditor struct {
	AdditionalAddressInfo *string `json:"additionalAddressInfo,omitempty"`
	City                  *string `json:"city,omitempty"`
	CountryCode           *string `json:"countryCode,omitempty"`
	HouseNumber           *string `json:"houseNumber,omitempty"`
	Iban                  *string `json:"iban,omitempty"`
	ID                    *string `json:"id,omitempty"`
	Name                  *string `json:"name,omitempty"`
	ReferenceParty        *string `json:"referenceParty,omitempty"`
	ReferencePartyID      *string `json:"referencePartyId,omitempty"`
	Street                *string `json:"street,omitempty"`
	Zip                   *string `json:"zip,omitempty"`
}

// NewCreditor constructs a new Creditor
func NewCreditor() *Creditor {
	return &Creditor{}
}
