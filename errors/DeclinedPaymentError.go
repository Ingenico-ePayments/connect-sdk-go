package errors

import (
	goerr "errors"
	"fmt"
	"strconv"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"
)

// DeclinedPaymentError represents an error response from a create payment call
type DeclinedPaymentError struct {
	errorMessage  string
	statusCode    int
	responseBody  string
	errorResponse *payment.ErrorResponse
}

// Message returns the error message
func (dpe *DeclinedPaymentError) Message() string {
	return dpe.errorMessage
}

// StatusCode returns the status code
func (dpe *DeclinedPaymentError) StatusCode() int {
	return dpe.statusCode
}

// ResponseBody returns the response body
func (dpe *DeclinedPaymentError) ResponseBody() string {
	return dpe.responseBody
}

// ErrorID returns the error id
func (dpe *DeclinedPaymentError) ErrorID() string {
	if dpe.errorResponse.ErrorID == nil {
		return ""
	}
	return *dpe.errorResponse.ErrorID
}

// Errors returns a slice of underlying errors
func (dpe *DeclinedPaymentError) Errors() []errors.APIError {
	// Return a clone instead of the original slice - immutability insurance
	if dpe.errorResponse.Errors == nil {
		return []errors.APIError{}
	}
	return append([]errors.APIError{}, *dpe.errorResponse.Errors...)
}

// PaymentResult returns the payment creation result
func (dpe *DeclinedPaymentError) PaymentResult() *payment.CreateResult {
	return dpe.errorResponse.PaymentResult
}

// String implements the Stringer interface
// Format: 'errorMessage; statusCode=; responseBody='
func (dpe *DeclinedPaymentError) String() string {
	list := dpe.errorMessage

	if dpe.statusCode > 0 {
		list = list + "; statusCode=" + strconv.Itoa(dpe.statusCode)
	}
	if len(dpe.responseBody) != 0 {
		list = list + "; responseBody='" + dpe.responseBody + "'"
	}

	return list
}

// Error implements the Error interface
func (dpe *DeclinedPaymentError) Error() string {
	return dpe.String()
}

// NewDeclinedPaymentError creates a DeclinedTransactionError with the given statusCode, responseBody and errorResponse
func NewDeclinedPaymentError(statusCode int, responseBody string, errorResponse *payment.ErrorResponse) (*DeclinedPaymentError, error) {
	if errorResponse.PaymentResult.Payment.ID == nil || errorResponse.PaymentResult.Payment.StatusOutput == nil {
		return nil, goerr.New("Cannot get payment")
	}
	errorMessage := fmt.Sprintf("declined payment '%s' with status '%s'",
		*errorResponse.PaymentResult.Payment.ID,
		*errorResponse.PaymentResult.Payment.Status)

	return &DeclinedPaymentError{errorMessage, statusCode, responseBody, errorResponse}, nil
}
