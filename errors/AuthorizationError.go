package errors

import (
	"strconv"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"
)

// AuthorizationError represents an error response from the GlobalCollect platform when authorization failed.
type AuthorizationError struct {
	errorMessage string
	statusCode   int
	responseBody string
	errorID      string
	errors       []errors.APIError
}

// Message returns the error message
func (ae *AuthorizationError) Message() string {
	return ae.errorMessage
}

// StatusCode returns the status code
func (ae *AuthorizationError) StatusCode() int {
	return ae.statusCode
}

// ResponseBody returns the response body
func (ae *AuthorizationError) ResponseBody() string {
	return ae.responseBody
}

// ErrorID returns the error id
func (ae *AuthorizationError) ErrorID() string {
	return ae.errorID
}

// Errors returns a slice of underlying errors
func (ae *AuthorizationError) Errors() []errors.APIError {
	// Return a clone instead of the original slice - immutability insurance
	return append([]errors.APIError{}, ae.errors...)
}

// String implements the Stringer interface
// Format: 'errorMessage; statusCode=; responseBody='
func (ae *AuthorizationError) String() (list string) {
	list = ae.errorMessage

	if ae.statusCode > 0 {
		list = list + "; statusCode=" + strconv.Itoa(ae.statusCode)
	}
	if len(ae.responseBody) != 0 {
		list = list + "; responseBody='" + ae.responseBody + "'"
	}

	return
}

// Error implements the Error interface
func (ae *AuthorizationError) Error() string {
	return ae.String()
}

// NewAuthorizationError creates an AuthorizationError with the given statusCode, responseBody, errorID and errors
func NewAuthorizationError(statusCode int, responseBody, errorID string, errors []errors.APIError) (*AuthorizationError, error) {
	return &AuthorizationError{"the GlobalCollect platform returned an incorrect request error response", statusCode, responseBody, errorID, errors}, nil
}

// NewAuthorizationErrorVerbose creates an AuthorizationError with the given message, statusCode, responseBody, errorID and errors
func NewAuthorizationErrorVerbose(message string, statusCode int, responseBody, errorID string, errors []errors.APIError) (*AuthorizationError, error) {
	return &AuthorizationError{message, statusCode, responseBody, errorID, errors}, nil
}
