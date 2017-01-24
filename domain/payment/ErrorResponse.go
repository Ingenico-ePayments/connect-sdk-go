// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"

// ErrorResponse represents class PaymentErrorResponse
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentErrorResponse
type ErrorResponse struct {
	ErrorID       *string            `json:"errorId,omitempty"`
	Errors        *[]errors.APIError `json:"errors,omitempty"`
	PaymentResult *CreateResult      `json:"paymentResult,omitempty"`
}

// NewErrorResponse constructs a new ErrorResponse
func NewErrorResponse() *ErrorResponse {
	return &ErrorResponse{}
}
