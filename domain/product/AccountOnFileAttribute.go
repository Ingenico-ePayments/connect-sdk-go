// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package product

// AccountOnFileAttribute represents class AccountOnFileAttribute
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
