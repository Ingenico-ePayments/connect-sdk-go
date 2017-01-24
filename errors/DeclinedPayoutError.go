package errors

import (
	goerr "errors"
	"fmt"
	"strconv"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payout"
)

// DeclinedPayoutError represents an error response from a create payout call
type DeclinedPayoutError struct {
	errorMessage  string
	statusCode    int
	responseBody  string
	errorResponse *payout.ErrorResponse
}

// Message returns the error message
func (dpe *DeclinedPayoutError) Message() string {
	return dpe.errorMessage
}

// StatusCode returns the status code
func (dpe *DeclinedPayoutError) StatusCode() int {
	return dpe.statusCode
}

// ResponseBody returns the response body
func (dpe *DeclinedPayoutError) ResponseBody() string {
	return dpe.responseBody
}

// ErrorID returns the error id
func (dpe *DeclinedPayoutError) ErrorID() string {
	if dpe.errorResponse.ErrorID == nil {
		return ""
	}
	return *dpe.errorResponse.ErrorID
}

// Errors returns a slice of underlying errors
func (dpe *DeclinedPayoutError) Errors() []errors.APIError {
	// Return a clone instead of the original slice - immutability insurance
	if dpe.errorResponse.Errors == nil {
		return []errors.APIError{}
	}
	return append([]errors.APIError{}, *dpe.errorResponse.Errors...)
}

// PayoutResult returns the payout result
func (dpe *DeclinedPayoutError) PayoutResult() *payout.Result {
	return dpe.errorResponse.PayoutResult
}

// String implements the Stringer interface
// Format: 'errorMessage; statusCode=; responseBody='
func (dpe *DeclinedPayoutError) String() string {
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
func (dpe *DeclinedPayoutError) Error() string {
	return dpe.String()
}

// NewDeclinedPayoutError creates a DeclinedPayoutError with the given statusCode, responseBody and errorResponse
func NewDeclinedPayoutError(statusCode int, responseBody string, errorResponse *payout.ErrorResponse) (*DeclinedPayoutError, error) {
	if errorResponse.PayoutResult.ID == nil || errorResponse.PayoutResult.Status == nil {
		return nil, goerr.New("Cannot get payout")
	}
	errorMessage := fmt.Sprintf("declined payout '%s' with status '%s'",
		*errorResponse.PayoutResult.ID,
		*errorResponse.PayoutResult.Status)

	return &DeclinedPayoutError{errorMessage, statusCode, responseBody, errorResponse}, nil
}
