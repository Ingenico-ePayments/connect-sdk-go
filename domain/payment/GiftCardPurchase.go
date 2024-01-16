// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// GiftCardPurchase represents class GiftCardPurchase
type GiftCardPurchase struct {
	AmountOfMoney     *definitions.AmountOfMoney `json:"amountOfMoney,omitempty"`
	NumberOfGiftCards *int32                     `json:"numberOfGiftCards,omitempty"`
}

// NewGiftCardPurchase constructs a new GiftCardPurchase
func NewGiftCardPurchase() *GiftCardPurchase {
	return &GiftCardPurchase{}
}
