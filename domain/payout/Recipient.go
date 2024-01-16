// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payout

// Recipient represents class PayoutRecipient
type Recipient struct {
	FirstName     *string `json:"firstName,omitempty"`
	Surname       *string `json:"surname,omitempty"`
	SurnamePrefix *string `json:"surnamePrefix,omitempty"`
}

// NewRecipient constructs a new Recipient
func NewRecipient() *Recipient {
	return &Recipient{}
}
