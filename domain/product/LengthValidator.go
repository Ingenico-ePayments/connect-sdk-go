// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package product

// LengthValidator represents class LengthValidator
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_LengthValidator
type LengthValidator struct {
	MaxLength *int32 `json:"maxLength,omitempty"`
	MinLength *int32 `json:"minLength,omitempty"`
}

// NewLengthValidator constructs a new LengthValidator
func NewLengthValidator() *LengthValidator {
	return &LengthValidator{}
}
