// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package token

// ContactDetailsToken represents class ContactDetailsToken
type ContactDetailsToken struct {
	EmailAddress     *string `json:"emailAddress,omitempty"`
	EmailMessageType *string `json:"emailMessageType,omitempty"`
}

// NewContactDetailsToken constructs a new ContactDetailsToken
func NewContactDetailsToken() *ContactDetailsToken {
	return &ContactDetailsToken{}
}
