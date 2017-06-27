// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// ContactDetails represents class ContactDetails
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
