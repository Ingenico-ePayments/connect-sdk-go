// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package product

// RegularExpressionValidator represents class RegularExpressionValidator
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_RegularExpressionValidator
type RegularExpressionValidator struct {
	RegularExpression *string `json:"regularExpression,omitempty"`
}

// NewRegularExpressionValidator constructs a new RegularExpressionValidator
func NewRegularExpressionValidator() *RegularExpressionValidator {
	return &RegularExpressionValidator{}
}
