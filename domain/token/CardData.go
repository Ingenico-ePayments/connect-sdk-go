// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package token

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// CardData represents class TokenCardData
type CardData struct {
	CardWithoutCvv       *definitions.CardWithoutCvv `json:"cardWithoutCvv,omitempty"`
	FirstTransactionDate *string                     `json:"firstTransactionDate,omitempty"`
	ProviderReference    *string                     `json:"providerReference,omitempty"`
}

// NewCardData constructs a new CardData
func NewCardData() *CardData {
	return &CardData{}
}
