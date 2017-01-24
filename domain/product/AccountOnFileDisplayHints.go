// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package product

// AccountOnFileDisplayHints represents class AccountOnFileDisplayHints
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_AccountOnFileDisplayHints
type AccountOnFileDisplayHints struct {
	LabelTemplate *[]LabelTemplateElement `json:"labelTemplate,omitempty"`
	Logo          *string                 `json:"logo,omitempty"`
}

// NewAccountOnFileDisplayHints constructs a new AccountOnFileDisplayHints
func NewAccountOnFileDisplayHints() *AccountOnFileDisplayHints {
	return &AccountOnFileDisplayHints{}
}
