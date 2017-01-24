// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package product

// PaymentProductFieldDisplayHints represents class PaymentProductFieldDisplayHints
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentProductFieldDisplayHints
type PaymentProductFieldDisplayHints struct {
	AlwaysShow         *bool                           `json:"alwaysShow,omitempty"`
	DisplayOrder       *int32                          `json:"displayOrder,omitempty"`
	FormElement        *PaymentProductFieldFormElement `json:"formElement,omitempty"`
	Label              *string                         `json:"label,omitempty"`
	Mask               *string                         `json:"mask,omitempty"`
	Obfuscate          *bool                           `json:"obfuscate,omitempty"`
	PlaceholderLabel   *string                         `json:"placeholderLabel,omitempty"`
	PreferredInputType *string                         `json:"preferredInputType,omitempty"`
	Tooltip            *PaymentProductFieldTooltip     `json:"tooltip,omitempty"`
}

// NewPaymentProductFieldDisplayHints constructs a new PaymentProductFieldDisplayHints
func NewPaymentProductFieldDisplayHints() *PaymentProductFieldDisplayHints {
	return &PaymentProductFieldDisplayHints{}
}
