// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package services

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// PaymentContext represents class PaymentContext
type PaymentContext struct {
	AmountOfMoney *definitions.AmountOfMoney `json:"amountOfMoney,omitempty"`
	CountryCode   *string                    `json:"countryCode,omitempty"`
	IsRecurring   *bool                      `json:"isRecurring,omitempty"`
}

// NewPaymentContext constructs a new PaymentContext
func NewPaymentContext() *PaymentContext {
	return &PaymentContext{}
}
