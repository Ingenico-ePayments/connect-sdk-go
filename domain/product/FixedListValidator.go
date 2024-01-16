// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package product

// FixedListValidator represents class FixedListValidator
type FixedListValidator struct {
	AllowedValues *[]string `json:"allowedValues,omitempty"`
}

// NewFixedListValidator constructs a new FixedListValidator
func NewFixedListValidator() *FixedListValidator {
	return &FixedListValidator{}
}
