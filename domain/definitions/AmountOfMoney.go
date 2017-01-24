// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package definitions

// AmountOfMoney represents class AmountOfMoney
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_AmountOfMoney
type AmountOfMoney struct {
	Amount       *int64  `json:"amount,omitempty"`
	CurrencyCode *string `json:"currencyCode,omitempty"`
}

// NewAmountOfMoney constructs a new AmountOfMoney
func NewAmountOfMoney() *AmountOfMoney {
	return &AmountOfMoney{}
}
