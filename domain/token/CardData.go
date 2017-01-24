// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package token

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// CardData represents class TokenCardData
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_TokenCardData
type CardData struct {
	CardWithoutCvv       *definitions.CardWithoutCvv `json:"cardWithoutCvv,omitempty"`
	FirstTransactionDate *string                     `json:"firstTransactionDate,omitempty"`
	ProviderReference    *string                     `json:"providerReference,omitempty"`
}

// NewCardData constructs a new CardData
func NewCardData() *CardData {
	return &CardData{}
}
