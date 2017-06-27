// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

// BankAccount represents class BankAccount
type BankAccount struct {
	AccountHolderName *string `json:"accountHolderName,omitempty"`
}

// NewBankAccount constructs a new BankAccount
func NewBankAccount() *BankAccount {
	return &BankAccount{}
}
