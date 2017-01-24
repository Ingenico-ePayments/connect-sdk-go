// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package product

// FixedListValidator represents class FixedListValidator
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_FixedListValidator
type FixedListValidator struct {
	AllowedValues *[]string `json:"allowedValues,omitempty"`
}

// NewFixedListValidator constructs a new FixedListValidator
func NewFixedListValidator() *FixedListValidator {
	return &FixedListValidator{}
}
