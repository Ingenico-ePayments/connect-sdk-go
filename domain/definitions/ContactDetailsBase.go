// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package definitions

// ContactDetailsBase represents class ContactDetailsBase
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_ContactDetailsBase
type ContactDetailsBase struct {
	EmailAddress     *string `json:"emailAddress,omitempty"`
	EmailMessageType *string `json:"emailMessageType,omitempty"`
}

// NewContactDetailsBase constructs a new ContactDetailsBase
func NewContactDetailsBase() *ContactDetailsBase {
	return &ContactDetailsBase{}
}
