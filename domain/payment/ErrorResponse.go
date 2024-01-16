// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"

// ErrorResponse represents class PaymentErrorResponse
type ErrorResponse struct {
	ErrorID       *string            `json:"errorId,omitempty"`
	Errors        *[]errors.APIError `json:"errors,omitempty"`
	PaymentResult *CreateResult      `json:"paymentResult,omitempty"`
}

// NewErrorResponse constructs a new ErrorResponse
func NewErrorResponse() *ErrorResponse {
	return &ErrorResponse{}
}
