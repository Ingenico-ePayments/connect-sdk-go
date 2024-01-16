// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// LoanRecipient represents class LoanRecipient
//
// Deprecated: No replacement
type LoanRecipient struct {
	// Deprecated: No replacement
	AccountNumber *string `json:"accountNumber,omitempty"`
	// Deprecated: No replacement
	DateOfBirth   *string `json:"dateOfBirth,omitempty"`
	// Deprecated: No replacement
	PartialPan    *string `json:"partialPan,omitempty"`
	// Deprecated: No replacement
	Surname       *string `json:"surname,omitempty"`
	// Deprecated: No replacement
	Zip           *string `json:"zip,omitempty"`
}

// NewLoanRecipient constructs a new LoanRecipient
func NewLoanRecipient() *LoanRecipient {
	return &LoanRecipient{}
}
