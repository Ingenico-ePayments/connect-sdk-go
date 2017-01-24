package errors

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"

// APIError represents an error response from the GlobalCollect platform which contains an ID and a list of errors.
type APIError interface {
	// Implements the error interface
	Error() string

	// Gets the raw response body that was returned by the GlobalConnect platform.
	Message() string

	// Gets the HTTP status code that was returned by the GlobalCollect platform.
	StatusCode() int

	// Gets the raw response body that was returned by the GlobalConnect platform.
	ResponseBody() string

	/// Gets the error identifier received from the GlobalCollect platform if available.
	ErrorID() string

	// Gets the error list received from the GlobalCollect if available. Never nil.
	Errors() []errors.APIError
}
