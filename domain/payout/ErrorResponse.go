// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payout

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"

// ErrorResponse represents class PayoutErrorResponse
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PayoutErrorResponse
type ErrorResponse struct {
	ErrorID      *string            `json:"errorId,omitempty"`
	Errors       *[]errors.APIError `json:"errors,omitempty"`
	PayoutResult *Result            `json:"payoutResult,omitempty"`
}

// NewErrorResponse constructs a new ErrorResponse
func NewErrorResponse() *ErrorResponse {
	return &ErrorResponse{}
}
