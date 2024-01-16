// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package product

// RegularExpressionValidator represents class RegularExpressionValidator
type RegularExpressionValidator struct {
	RegularExpression *string `json:"regularExpression,omitempty"`
}

// NewRegularExpressionValidator constructs a new RegularExpressionValidator
func NewRegularExpressionValidator() *RegularExpressionValidator {
	return &RegularExpressionValidator{}
}
