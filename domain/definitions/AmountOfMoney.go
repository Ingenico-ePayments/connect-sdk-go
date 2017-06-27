// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

// AmountOfMoney represents class AmountOfMoney
type AmountOfMoney struct {
	Amount       *int64  `json:"amount,omitempty"`
	CurrencyCode *string `json:"currencyCode,omitempty"`
}

// NewAmountOfMoney constructs a new AmountOfMoney
func NewAmountOfMoney() *AmountOfMoney {
	return &AmountOfMoney{}
}
