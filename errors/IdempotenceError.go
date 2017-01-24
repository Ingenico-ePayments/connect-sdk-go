package errors

import (
	"strconv"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"
)

// IdempotenceError represents an error response from the GlobalCollect platform when an idempotent request failed because the first request has not finished yet.
type IdempotenceError struct {
	idempotenceKey              string
	idempotenceRequestTimestamp *int64

	errorMessage string
	statusCode   int
	responseBody string
	errorID      string
	errors       []errors.APIError
}

// Message returns the error message
func (ie *IdempotenceError) Message() string {
	return ie.errorMessage
}

// StatusCode returns the status code
func (ie *IdempotenceError) StatusCode() int {
	return ie.statusCode
}

// ResponseBody returns the response body
func (ie *IdempotenceError) ResponseBody() string {
	return ie.responseBody
}

// ErrorID returns the error id
func (ie *IdempotenceError) ErrorID() string {
	return ie.errorID
}

// Errors returns a slice of underlying errors
func (ie *IdempotenceError) Errors() []errors.APIError {
	// Return a clone instead of the original slice - immutability insurance
	return append([]errors.APIError{}, ie.errors...)
}

// IdempotenceKey returns the idempotence key used
func (ie *IdempotenceError) IdempotenceKey() string {
	return ie.idempotenceKey
}

// IdempotenceRequestTimestamp returns the timestamp of the request
func (ie *IdempotenceError) IdempotenceRequestTimestamp() *int64 {
	var timestamp *int64

	if ie.idempotenceRequestTimestamp == nil {
		timestamp = nil
	} else {
		timestamp = new(int64)

		*timestamp = *ie.idempotenceRequestTimestamp
	}

	return timestamp
}

// String implements the Stringer interface
// Format: 'errorMessage; statusCode=; responseBody='
func (ie *IdempotenceError) String() string {
	list := ie.errorMessage

	if ie.statusCode > 0 {
		list = list + "; statusCode=" + strconv.Itoa(ie.statusCode)
	}
	if len(ie.responseBody) != 0 {
		list = list + "; responseBody='" + ie.responseBody + "'"
	}

	return list
}

// Error implements the Error interface
func (ie *IdempotenceError) Error() string {
	return ie.String()
}

// NewIdempotenceError creates an IdempotenceError with the given idempotenceKey, idempotenceRequestTimestamp, statusCode, responseBody, errorID and errors
func NewIdempotenceError(idempotenceKey string, idempotenceRequestTimestamp *int64, statusCode int, responseBody, errorID string, errors []errors.APIError) (*IdempotenceError, error) {
	return &IdempotenceError{idempotenceKey, idempotenceRequestTimestamp, "the GlobalCollect platform returned an incorrect request error response", statusCode, responseBody, errorID, errors}, nil
}

// NewIdempotenceErrorVerbose creates an IdempotenceError with the given idempotenceKey, idempotenceRequestTimestamp, message, statusCode, responseBody, errorID and errors
func NewIdempotenceErrorVerbose(idempotenceKey string, idempotenceRequestTimestamp *int64, message string, statusCode int, responseBody, errorID string, errors []errors.APIError) (*IdempotenceError, error) {
	return &IdempotenceError{idempotenceKey, idempotenceRequestTimestamp, message, statusCode, responseBody, errorID, errors}, nil
}
