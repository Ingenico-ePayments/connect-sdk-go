// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// LoanRecipient represents class LoanRecipient
type LoanRecipient struct {
	AccountNumber *string `json:"accountNumber,omitempty"`
	DateOfBirth   *string `json:"dateOfBirth,omitempty"`
	PartialPan    *string `json:"partialPan,omitempty"`
	Surname       *string `json:"surname,omitempty"`
	Zip           *string `json:"zip,omitempty"`
}

// NewLoanRecipient constructs a new LoanRecipient
func NewLoanRecipient() *LoanRecipient {
	return &LoanRecipient{}
}
