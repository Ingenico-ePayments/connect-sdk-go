// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package refund

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"

// ErrorResponse represents class RefundErrorResponse
type ErrorResponse struct {
	ErrorID      *string            `json:"errorId,omitempty"`
	Errors       *[]errors.APIError `json:"errors,omitempty"`
	RefundResult *Result            `json:"refundResult,omitempty"`
}

// NewErrorResponse constructs a new ErrorResponse
func NewErrorResponse() *ErrorResponse {
	return &ErrorResponse{}
}
