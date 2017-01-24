// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payout

import (
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"
)

// Result represents class PayoutResult
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PayoutResult
type Result struct {
	ID           *string                        `json:"id,omitempty"`
	PayoutOutput *payment.OrderOutput           `json:"payoutOutput,omitempty"`
	Status       *string                        `json:"status,omitempty"`
	StatusOutput *definitions.OrderStatusOutput `json:"statusOutput,omitempty"`
}

// NewResult constructs a new Result
func NewResult() *Result {
	return &Result{}
}
