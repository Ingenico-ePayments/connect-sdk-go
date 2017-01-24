// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package product

// RangeValidator represents class RangeValidator
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_RangeValidator
type RangeValidator struct {
	MaxValue *int32 `json:"maxValue,omitempty"`
	MinValue *int32 `json:"minValue,omitempty"`
}

// NewRangeValidator constructs a new RangeValidator
func NewRangeValidator() *RangeValidator {
	return &RangeValidator{}
}
