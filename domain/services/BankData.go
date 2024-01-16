// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package services

// BankData represents class BankData
type BankData struct {
	NewBankName              *string `json:"newBankName,omitempty"`
	ReformattedAccountNumber *string `json:"reformattedAccountNumber,omitempty"`
	ReformattedBankCode      *string `json:"reformattedBankCode,omitempty"`
	ReformattedBranchCode    *string `json:"reformattedBranchCode,omitempty"`
}

// NewBankData constructs a new BankData
func NewBankData() *BankData {
	return &BankData{}
}
