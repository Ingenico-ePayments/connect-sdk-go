// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package definitions

// ContactDetailsBase represents class ContactDetailsBase
type ContactDetailsBase struct {
	EmailAddress     *string `json:"emailAddress,omitempty"`
	EmailMessageType *string `json:"emailMessageType,omitempty"`
}

// NewContactDetailsBase constructs a new ContactDetailsBase
func NewContactDetailsBase() *ContactDetailsBase {
	return &ContactDetailsBase{}
}
