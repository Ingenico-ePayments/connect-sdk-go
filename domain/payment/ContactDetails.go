// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// ContactDetails represents class ContactDetails
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_ContactDetails
type ContactDetails struct {
	EmailAddress     *string `json:"emailAddress,omitempty"`
	EmailMessageType *string `json:"emailMessageType,omitempty"`
	FaxNumber        *string `json:"faxNumber,omitempty"`
	PhoneNumber      *string `json:"phoneNumber,omitempty"`
}

// NewContactDetails constructs a new ContactDetails
func NewContactDetails() *ContactDetails {
	return &ContactDetails{}
}
