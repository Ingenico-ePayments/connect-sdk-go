// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

// BankAccountIban represents class BankAccountIban
type BankAccountIban struct {
	AccountHolderName *string `json:"accountHolderName,omitempty"`
	Iban              *string `json:"iban,omitempty"`
}

// NewBankAccountIban constructs a new BankAccountIban
func NewBankAccountIban() *BankAccountIban {
	return &BankAccountIban{}
}
