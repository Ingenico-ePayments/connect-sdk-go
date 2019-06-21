package communicator

import "io"

// Marshaller is the interface used to marshal and unmarshal Ingenico ePayments platform request
// and response objects to and from JSON. Thread-safe.
type Marshaller interface {
	/// Marshal a request object to a JSON string.
	Marshal(interface{}) (string, error)

	/// Unmarshal a JSON string to a response object.
	Unmarshal(string, interface{}) error

	/// Unmarshal a io.Reader to a response object.
	UnmarshalFromReader(io.Reader, interface{}) error
}
