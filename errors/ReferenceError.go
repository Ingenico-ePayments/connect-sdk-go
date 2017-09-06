package errors

import (
	"strconv"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"
)

// ReferenceError represents an error response from the Ingenico ePayments platform when a non-existing or removed object is trying to be accessed.
type ReferenceError struct {
	errorMessage string
	statusCode   int
	responseBody string
	errorID      string
	errors       []errors.APIError
}

// Message returns the error message
func (re *ReferenceError) Message() string {
	return re.errorMessage
}

// StatusCode returns the status code
func (re *ReferenceError) StatusCode() int {
	return re.statusCode
}

// ResponseBody returns the response body
func (re *ReferenceError) ResponseBody() string {
	return re.responseBody
}

// ErrorID returns the error id
func (re *ReferenceError) ErrorID() string {
	return re.errorID
}

// Errors returns a slice of underlying errors
func (re *ReferenceError) Errors() []errors.APIError {
	// Return a clone instead of the original slice - immutability insurance
	return append([]errors.APIError{}, re.errors...)
}

// String implements the Stringer interface
// Format: 'errorMessage; statusCode=; responseBody='
func (re *ReferenceError) String() string {
	list := re.errorMessage

	if re.statusCode > 0 {
		list = list + "; statusCode=" + strconv.Itoa(re.statusCode)
	}
	if len(re.responseBody) != 0 {
		list = list + "; responseBody='" + re.responseBody + "'"
	}

	return list
}

// Error implements the Error interface
func (re *ReferenceError) Error() string {
	return re.String()
}

// NewReferenceError creates a ReferenceError with the given statusCode, responseBody, errorID and errors
func NewReferenceError(statusCode int, responseBody, errorID string, errors []errors.APIError) (*ReferenceError, error) {
	return &ReferenceError{"the Ingenico ePayments platform returned an incorrect request error response", statusCode, responseBody, errorID, errors}, nil
}

// NewReferenceErrorVerbose creates a ReferenceError with the given message, statusCode, responseBody, errorID and errors
func NewReferenceErrorVerbose(message string, statusCode int, responseBody, errorID string, errors []errors.APIError) (*ReferenceError, error) {
	return &ReferenceError{message, statusCode, responseBody, errorID, errors}, nil
}
