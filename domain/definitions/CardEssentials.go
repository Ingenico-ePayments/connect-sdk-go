// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package definitions

// CardEssentials represents class CardEssentials
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_CardEssentials
type CardEssentials struct {
	CardNumber *string `json:"cardNumber,omitempty"`
	ExpiryDate *string `json:"expiryDate,omitempty"`
}

// NewCardEssentials constructs a new CardEssentials
func NewCardEssentials() *CardEssentials {
	return &CardEssentials{}
}
