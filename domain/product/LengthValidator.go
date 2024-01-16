// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package product

// LengthValidator represents class LengthValidator
type LengthValidator struct {
	MaxLength *int32 `json:"maxLength,omitempty"`
	MinLength *int32 `json:"minLength,omitempty"`
}

// NewLengthValidator constructs a new LengthValidator
func NewLengthValidator() *LengthValidator {
	return &LengthValidator{}
}
