// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package refund

// BankAccountBbanRefund represents class BankAccountBbanRefund
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_BankAccountBbanRefund
type BankAccountBbanRefund struct {
	AccountHolderName *string `json:"accountHolderName,omitempty"`
	AccountNumber     *string `json:"accountNumber,omitempty"`
	BankCity          *string `json:"bankCity,omitempty"`
	BankCode          *string `json:"bankCode,omitempty"`
	BankName          *string `json:"bankName,omitempty"`
	BranchCode        *string `json:"branchCode,omitempty"`
	CheckDigit        *string `json:"checkDigit,omitempty"`
	CountryCode       *string `json:"countryCode,omitempty"`
	SwiftCode         *string `json:"swiftCode,omitempty"`
}

// NewBankAccountBbanRefund constructs a new BankAccountBbanRefund
func NewBankAccountBbanRefund() *BankAccountBbanRefund {
	return &BankAccountBbanRefund{}
}
