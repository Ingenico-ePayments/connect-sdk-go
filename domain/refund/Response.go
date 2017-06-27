// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package refund

import (
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"
)

// Response represents class RefundResponse
type Response struct {
	ID           *string                        `json:"id,omitempty"`
	RefundOutput *payment.RefundOutput          `json:"refundOutput,omitempty"`
	Status       *string                        `json:"status,omitempty"`
	StatusOutput *definitions.OrderStatusOutput `json:"statusOutput,omitempty"`
}

// NewResponse constructs a new Response
func NewResponse() *Response {
	return &Response{}
}
