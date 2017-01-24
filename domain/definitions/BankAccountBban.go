// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package definitions

// BankAccountBban represents class BankAccountBban
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_BankAccountBban
type BankAccountBban struct {
	AccountHolderName *string `json:"accountHolderName,omitempty"`
	AccountNumber     *string `json:"accountNumber,omitempty"`
	BankCode          *string `json:"bankCode,omitempty"`
	BankName          *string `json:"bankName,omitempty"`
	BranchCode        *string `json:"branchCode,omitempty"`
	CheckDigit        *string `json:"checkDigit,omitempty"`
	CountryCode       *string `json:"countryCode,omitempty"`
}

// NewBankAccountBban constructs a new BankAccountBban
func NewBankAccountBban() *BankAccountBban {
	return &BankAccountBban{}
}
