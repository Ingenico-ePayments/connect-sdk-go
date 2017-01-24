// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package token

// Card represents class TokenCard
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_TokenCard
type Card struct {
	Alias    *string        `json:"alias,omitempty"`
	Customer *CustomerToken `json:"customer,omitempty"`
	Data     *CardData      `json:"data,omitempty"`
}

// NewCard constructs a new Card
func NewCard() *Card {
	return &Card{}
}
