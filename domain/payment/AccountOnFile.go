// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// AccountOnFile represents class PaymentAccountOnFile
type AccountOnFile struct {
	CreateDate                                    *string `json:"createDate,omitempty"`
	NumberOfCardOnFileCreationAttemptsLast24Hours *int32  `json:"numberOfCardOnFileCreationAttemptsLast24Hours,omitempty"`
}

// NewAccountOnFile constructs a new AccountOnFile
func NewAccountOnFile() *AccountOnFile {
	return &AccountOnFile{}
}
