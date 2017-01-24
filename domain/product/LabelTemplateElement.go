// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package product

// LabelTemplateElement represents class LabelTemplateElement
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_LabelTemplateElement
type LabelTemplateElement struct {
	AttributeKey *string `json:"attributeKey,omitempty"`
	Mask         *string `json:"mask,omitempty"`
}

// NewLabelTemplateElement constructs a new LabelTemplateElement
func NewLabelTemplateElement() *LabelTemplateElement {
	return &LabelTemplateElement{}
}
