// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// TrustlyBankAccount represents class TrustlyBankAccount
type TrustlyBankAccount struct {
	AccountLastDigits          *string `json:"accountLastDigits,omitempty"`
	BankName                   *string `json:"bankName,omitempty"`
	Clearinghouse              *string `json:"clearinghouse,omitempty"`
	PersonIdentificationNumber *string `json:"personIdentificationNumber,omitempty"`
}

// NewTrustlyBankAccount constructs a new TrustlyBankAccount
func NewTrustlyBankAccount() *TrustlyBankAccount {
	return &TrustlyBankAccount{}
}
