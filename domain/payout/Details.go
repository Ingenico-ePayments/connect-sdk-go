// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payout

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// Details represents class PayoutDetails
type Details struct {
	AmountOfMoney *definitions.AmountOfMoney `json:"amountOfMoney,omitempty"`
	Customer      *Customer                  `json:"customer,omitempty"`
	References    *References                `json:"references,omitempty"`
}

// NewDetails constructs a new Details
func NewDetails() *Details {
	return &Details{}
}
