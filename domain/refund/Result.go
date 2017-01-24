// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package refund

import (
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"
)

// Result represents class RefundResult
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_RefundResult
type Result struct {
	ID           *string                        `json:"id,omitempty"`
	RefundOutput *payment.RefundOutput          `json:"refundOutput,omitempty"`
	Status       *string                        `json:"status,omitempty"`
	StatusOutput *definitions.OrderStatusOutput `json:"statusOutput,omitempty"`
}

// NewResult constructs a new Result
func NewResult() *Result {
	return &Result{}
}
