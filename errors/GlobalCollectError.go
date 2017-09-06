package errors

import (
	"strconv"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"
)

// GlobalCollectError represents an error response from the Ingenico ePayments platform when something went wrong at the Ingenico ePayments platform or further downstream.
type GlobalCollectError struct {
	errorMessage string
	statusCode   int
	responseBody string
	errorID      string
	errors       []errors.APIError
}

// Message returns the error message
func (gce *GlobalCollectError) Message() string {
	return gce.errorMessage
}

// StatusCode returns the status code
func (gce *GlobalCollectError) StatusCode() int {
	return gce.statusCode
}

// ResponseBody returns the response body
func (gce *GlobalCollectError) ResponseBody() string {
	return gce.responseBody
}

// ErrorID returns the error id
func (gce *GlobalCollectError) ErrorID() string {
	return gce.errorID
}

// Errors returns a slice of underlying errors
func (gce *GlobalCollectError) Errors() []errors.APIError {
	// Return a clone instead of the original slice - immutability insurance
	return append([]errors.APIError{}, gce.errors...)
}

// String implements the Stringer ineterface
// Format: 'errorMessage; statusCode=; responseBody='
func (gce *GlobalCollectError) String() string {
	list := gce.errorMessage

	if gce.statusCode > 0 {
		list = list + "; statusCode=" + strconv.Itoa(gce.statusCode)
	}
	if len(gce.responseBody) != 0 {
		list = list + "; responseBody='" + gce.responseBody + "'"
	}

	return list
}

// Error implements the Error interface
func (gce *GlobalCollectError) Error() string {
	return gce.String()
}

// NewGlobalCollectError creates a GlobalCollectError with the given statusCode, responseBody, errorID and errors
func NewGlobalCollectError(statusCode int, responseBody, errorID string, errors []errors.APIError) (*GlobalCollectError, error) {
	return &GlobalCollectError{"the Ingenico ePayments platform returned an error response", statusCode, responseBody, errorID, errors}, nil
}

// NewGlobalCollectErrorVerbose creates a GlobalCollectError with the given message, statusCode, responseBody, errorID and errors
func NewGlobalCollectErrorVerbose(message string, statusCode int, responseBody, errorID string, errors []errors.APIError) (*GlobalCollectError, error) {
	return &GlobalCollectError{message, statusCode, responseBody, errorID, errors}, nil
}
