package errors

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"

// APIError represents an error response from the Ingenico ePayments platform which contains an ID and a list of errors.
type APIError interface {
	// Error implements the error interface
	Error() string

	// Message gets the raw response body that was returned by the Ingenico ePayments platform.
	Message() string

	// StatusCode gets the HTTP status code that was returned by the Ingenico ePayments platform.
	StatusCode() int

	// ResponseBody gets the raw response body that was returned by the Ingenico ePayments platform.
	ResponseBody() string

	// ErrorID gets the error identifier received from the Ingenico ePayments platform if available.
	ErrorID() string

	// Errors gets the error list received from the Ingenico ePayments platform if available. Never nil.
	Errors() []errors.APIError
}
