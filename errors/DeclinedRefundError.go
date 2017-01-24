package errors

import (
	goerr "errors"
	"fmt"
	"strconv"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/refund"
)

// DeclinedRefundError represents an error response from a refund call
type DeclinedRefundError struct {
	errorMessage  string
	statusCode    int
	responseBody  string
	errorResponse *refund.ErrorResponse
}

// Message returns the error message
func (dpe *DeclinedRefundError) Message() string {
	return dpe.errorMessage
}

// StatusCode returns the status code
func (dpe *DeclinedRefundError) StatusCode() int {
	return dpe.statusCode
}

// ResponseBody returns the response body
func (dpe *DeclinedRefundError) ResponseBody() string {
	return dpe.responseBody
}

// ErrorID returns the error id
func (dpe *DeclinedRefundError) ErrorID() string {
	if dpe.errorResponse.ErrorID == nil {
		return ""
	}
	return *dpe.errorResponse.ErrorID
}

// Errors returns a slice of underlying errors
func (dpe *DeclinedRefundError) Errors() []errors.APIError {
	// Return a clone instead of the original slice - immutability insurance
	if dpe.errorResponse.Errors == nil {
		return []errors.APIError{}
	}
	return append([]errors.APIError{}, *dpe.errorResponse.Errors...)
}

// RefundResult returns the refund result
func (dpe *DeclinedRefundError) RefundResult() *refund.Result {
	return dpe.errorResponse.RefundResult
}

// String implements the Stringer interface
// Format: 'errorMessage; statusCode=; responseBody='
func (dpe *DeclinedRefundError) String() string {
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
func (dpe *DeclinedRefundError) Error() string {
	return dpe.String()
}

// NewDeclinedRefundError creates a DeclinedRefundError with the given statusCode, responseBody and errorResponse
func NewDeclinedRefundError(statusCode int, responseBody string, errorResponse *refund.ErrorResponse) (*DeclinedRefundError, error) {
	if errorResponse.RefundResult.ID == nil || errorResponse.RefundResult.Status == nil {
		return nil, goerr.New("Cannot get refund")
	}
	errorMessage := fmt.Sprintf("declined refund '%s' with status '%s'",
		*errorResponse.RefundResult.ID,
		*errorResponse.RefundResult.Status)

	return &DeclinedRefundError{errorMessage, statusCode, responseBody, errorResponse}, nil
}
