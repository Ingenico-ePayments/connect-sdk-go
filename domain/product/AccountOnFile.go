// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package product

// AccountOnFile represents class AccountOnFile
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_AccountOnFile
type AccountOnFile struct {
	Attributes       *[]AccountOnFileAttribute  `json:"attributes,omitempty"`
	DisplayHints     *AccountOnFileDisplayHints `json:"displayHints,omitempty"`
	ID               *int32                     `json:"id,omitempty"`
	PaymentProductID *int32                     `json:"paymentProductId,omitempty"`
}

// NewAccountOnFile constructs a new AccountOnFile
func NewAccountOnFile() *AccountOnFile {
	return &AccountOnFile{}
}
