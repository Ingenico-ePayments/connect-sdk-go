// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package definitions

// CardEssentials represents class CardEssentials
type CardEssentials struct {
	CardNumber     *string `json:"cardNumber,omitempty"`
	CardholderName *string `json:"cardholderName,omitempty"`
	ExpiryDate     *string `json:"expiryDate,omitempty"`
}

// NewCardEssentials constructs a new CardEssentials
func NewCardEssentials() *CardEssentials {
	return &CardEssentials{}
}
