// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package token

// Card represents class TokenCard
type Card struct {
	Alias    *string        `json:"alias,omitempty"`
	Customer *CustomerToken `json:"customer,omitempty"`
	Data     *CardData      `json:"data,omitempty"`
}

// NewCard constructs a new Card
func NewCard() *Card {
	return &Card{}
}
