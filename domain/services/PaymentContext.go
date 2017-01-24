// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package services

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// PaymentContext represents class PaymentContext
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentContext
type PaymentContext struct {
	AmountOfMoney *definitions.AmountOfMoney `json:"amountOfMoney,omitempty"`
	CountryCode   *string                    `json:"countryCode,omitempty"`
	IsRecurring   *bool                      `json:"isRecurring,omitempty"`
}

// NewPaymentContext constructs a new PaymentContext
func NewPaymentContext() *PaymentContext {
	return &PaymentContext{}
}
