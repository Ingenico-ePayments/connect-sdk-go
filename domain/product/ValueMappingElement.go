// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package product

// ValueMappingElement represents class ValueMappingElement
type ValueMappingElement struct {
	DisplayElements *[]PaymentProductFieldDisplayElement `json:"displayElements,omitempty"`
	// Deprecated: use displayElement with ID 'displayName' instead.
	DisplayName     *string                              `json:"displayName,omitempty"`
	Value           *string                              `json:"value,omitempty"`
}

// NewValueMappingElement constructs a new ValueMappingElement
func NewValueMappingElement() *ValueMappingElement {
	return &ValueMappingElement{}
}
