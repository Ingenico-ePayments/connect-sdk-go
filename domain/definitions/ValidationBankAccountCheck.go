// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

// ValidationBankAccountCheck represents class ValidationBankAccountCheck
type ValidationBankAccountCheck struct {
	Code        *string `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
	Result      *string `json:"result,omitempty"`
}

// NewValidationBankAccountCheck constructs a new ValidationBankAccountCheck
func NewValidationBankAccountCheck() *ValidationBankAccountCheck {
	return &ValidationBankAccountCheck{}
}
