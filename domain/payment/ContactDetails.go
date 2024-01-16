// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// ContactDetails represents class ContactDetails
type ContactDetails struct {
	EmailAddress      *string `json:"emailAddress,omitempty"`
	EmailMessageType  *string `json:"emailMessageType,omitempty"`
	FaxNumber         *string `json:"faxNumber,omitempty"`
	MobilePhoneNumber *string `json:"mobilePhoneNumber,omitempty"`
	PhoneNumber       *string `json:"phoneNumber,omitempty"`
	WorkPhoneNumber   *string `json:"workPhoneNumber,omitempty"`
}

// NewContactDetails constructs a new ContactDetails
func NewContactDetails() *ContactDetails {
	return &ContactDetails{}
}
