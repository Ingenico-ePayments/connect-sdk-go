// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package product

// RangeValidator represents class RangeValidator
type RangeValidator struct {
	MaxValue *int32 `json:"maxValue,omitempty"`
	MinValue *int32 `json:"minValue,omitempty"`
}

// NewRangeValidator constructs a new RangeValidator
func NewRangeValidator() *RangeValidator {
	return &RangeValidator{}
}
