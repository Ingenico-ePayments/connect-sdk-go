// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package product

// ValueMappingElement represents class ValueMappingElement
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_ValueMappingElement
type ValueMappingElement struct {
	DisplayName *string `json:"displayName,omitempty"`
	Value       *string `json:"value,omitempty"`
}

// NewValueMappingElement constructs a new ValueMappingElement
func NewValueMappingElement() *ValueMappingElement {
	return &ValueMappingElement{}
}
