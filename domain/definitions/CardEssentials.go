// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

// CardEssentials represents class CardEssentials
type CardEssentials struct {
	CardNumber *string `json:"cardNumber,omitempty"`
	ExpiryDate *string `json:"expiryDate,omitempty"`
}

// NewCardEssentials constructs a new CardEssentials
func NewCardEssentials() *CardEssentials {
	return &CardEssentials{}
}
