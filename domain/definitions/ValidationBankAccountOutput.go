// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package definitions

// ValidationBankAccountOutput represents class ValidationBankAccountOutput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_ValidationBankAccountOutput
type ValidationBankAccountOutput struct {
	Checks                   *[]ValidationBankAccountCheck `json:"checks,omitempty"`
	NewBankName              *string                       `json:"newBankName,omitempty"`
	ReformattedAccountNumber *string                       `json:"reformattedAccountNumber,omitempty"`
	ReformattedBankCode      *string                       `json:"reformattedBankCode,omitempty"`
	ReformattedBranchCode    *string                       `json:"reformattedBranchCode,omitempty"`
}

// NewValidationBankAccountOutput constructs a new ValidationBankAccountOutput
func NewValidationBankAccountOutput() *ValidationBankAccountOutput {
	return &ValidationBankAccountOutput{}
}
