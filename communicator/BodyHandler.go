package communicator

import "io"
import "github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"

// BodyHandlerFunc is a handler function for an incoming body stream
type BodyHandlerFunc func(headers []communication.Header, reader io.Reader) error

// BodyHandler is a handler for an incoming body stream
type BodyHandler interface {
	Handle(headers []communication.Header, reader io.Reader) error
}

// Handle calls f(statusCode, headers, reader)
func (f BodyHandlerFunc) Handle(headers []communication.Header, reader io.Reader) error {
	return f(headers, reader)
}
