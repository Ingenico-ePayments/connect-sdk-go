// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payout

import (
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"
)

// Response represents class PayoutResponse
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PayoutResponse
type Response struct {
	ID           *string                        `json:"id,omitempty"`
	PayoutOutput *payment.OrderOutput           `json:"payoutOutput,omitempty"`
	Status       *string                        `json:"status,omitempty"`
	StatusOutput *definitions.OrderStatusOutput `json:"statusOutput,omitempty"`
}

// NewResponse constructs a new Response
func NewResponse() *Response {
	return &Response{}
}
