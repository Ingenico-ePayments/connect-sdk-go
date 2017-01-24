package errors

import (
	"strconv"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"
)

// ValidateError represents an error response from the GlobalCollect platform when validation of requests failed.
type ValidateError struct {
	errorMessage string
	statusCode   int
	responseBody string
	errorID      string
	errors       []errors.APIError
}

// Message returns the error message
func (ve *ValidateError) Message() string {
	return ve.errorMessage
}

// StatusCode returns the status code
func (ve *ValidateError) StatusCode() int {
	return ve.statusCode
}

// ResponseBody returns the response body
func (ve *ValidateError) ResponseBody() string {
	return ve.responseBody
}

// ErrorID returns the error id
func (ve *ValidateError) ErrorID() string {
	return ve.errorID
}

// Errors returns a slice of underlying errors
func (ve *ValidateError) Errors() []errors.APIError {
	// Return a clone instead of the original slice - immutability insurance
	return append([]errors.APIError{}, ve.errors...)
}

// String implements the Stringer interface
// Format: 'errorMessage; statusCode=; responseBody='
func (ve *ValidateError) String() string {
	list := ve.errorMessage

	if ve.statusCode > 0 {
		list = list + "; statusCode=" + strconv.Itoa(ve.statusCode)
	}
	if len(ve.responseBody) != 0 {
		list = list + "; responseBody='" + ve.responseBody + "'"
	}

	return list
}

// Error implements the Error interface
func (ve *ValidateError) Error() string {
	return ve.String()
}

// NewValidateError creates a ValidateError with the given statusCode, responseBody, errorID and errors
func NewValidateError(statusCode int, responseBody, errorID string, errors []errors.APIError) (*ValidateError, error) {
	return &ValidateError{"the GlobalCollect platform returned an incorrect request error response", statusCode, responseBody, errorID, errors}, nil
}

// NewValidateErrorVerbose creates a ValidateError with the given message, statusCode, responseBody, errorID and errors
func NewValidateErrorVerbose(message string, statusCode int, responseBody, errorID string, errors []errors.APIError) (*ValidateError, error) {
	return &ValidateError{message, statusCode, responseBody, errorID, errors}, nil
}
