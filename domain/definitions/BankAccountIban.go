// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package definitions

// BankAccountIban represents class BankAccountIban
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_BankAccountIban
type BankAccountIban struct {
	AccountHolderName *string `json:"accountHolderName,omitempty"`
	Iban              *string `json:"iban,omitempty"`
}

// NewBankAccountIban constructs a new BankAccountIban
func NewBankAccountIban() *BankAccountIban {
	return &BankAccountIban{}
}
