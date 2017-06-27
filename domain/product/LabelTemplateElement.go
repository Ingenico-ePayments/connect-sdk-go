// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package product

// LabelTemplateElement represents class LabelTemplateElement
type LabelTemplateElement struct {
	AttributeKey *string `json:"attributeKey,omitempty"`
	Mask         *string `json:"mask,omitempty"`
}

// NewLabelTemplateElement constructs a new LabelTemplateElement
func NewLabelTemplateElement() *LabelTemplateElement {
	return &LabelTemplateElement{}
}
