// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package token

// ContactDetailsToken represents class ContactDetailsToken
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_ContactDetailsToken
type ContactDetailsToken struct {
	EmailAddress     *string `json:"emailAddress,omitempty"`
	EmailMessageType *string `json:"emailMessageType,omitempty"`
}

// NewContactDetailsToken constructs a new ContactDetailsToken
func NewContactDetailsToken() *ContactDetailsToken {
	return &ContactDetailsToken{}
}
