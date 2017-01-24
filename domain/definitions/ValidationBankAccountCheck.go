// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package definitions

// ValidationBankAccountCheck represents class ValidationBankAccountCheck
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_ValidationBankAccountCheck
type ValidationBankAccountCheck struct {
	Code        *string `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
	Result      *string `json:"result,omitempty"`
}

// NewValidationBankAccountCheck constructs a new ValidationBankAccountCheck
func NewValidationBankAccountCheck() *ValidationBankAccountCheck {
	return &ValidationBankAccountCheck{}
}
