// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package product

// AccountOnFileAttribute represents class AccountOnFileAttribute
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_AccountOnFileAttribute
type AccountOnFileAttribute struct {
	Key             *string `json:"key,omitempty"`
	MustWriteReason *string `json:"mustWriteReason,omitempty"`
	Status          *string `json:"status,omitempty"`
	Value           *string `json:"value,omitempty"`
}

// NewAccountOnFileAttribute constructs a new AccountOnFileAttribute
func NewAccountOnFileAttribute() *AccountOnFileAttribute {
	return &AccountOnFileAttribute{}
}
