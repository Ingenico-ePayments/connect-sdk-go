// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package definitions

// BankAccount represents class BankAccount
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_BankAccount
type BankAccount struct {
	AccountHolderName *string `json:"accountHolderName,omitempty"`
}

// NewBankAccount constructs a new BankAccount
func NewBankAccount() *BankAccount {
	return &BankAccount{}
}
