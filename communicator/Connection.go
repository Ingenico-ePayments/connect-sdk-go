package communicator

import (
	"net/url"
	"time"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"
	"github.com/Ingenico-ePayments/connect-sdk-go/logging"
)

// Connection represents a pooled connection to the Ingenico ePayments platform server.
// Instead of setting up a new HTTP connection for each request, this
// connection uses a pool of HTTP connections.
// Thread-safe
type Connection interface {

	// CloseIdleConnection closes all HTTP connections that have been idle for the specified time. This should also include
	// all expired HTTP connections.
	// timespan represents the time spent idle
	CloseIdleConnections(time time.Duration)

	// CloseExpiredConnections closes all expired HTTP connections.
	CloseExpiredConnections()

	//IMPLEMENTATION Connection INTERFACE

	// Get sends a GET request to the Ingenico ePayments platform and return the response.
	Get(resourceURI url.URL, requestHeaders []communication.Header) (*communication.Response, error)

	// Delete sends a DELETE request to the Ingenico ePayments platform and return the response.
	Delete(resourceURI url.URL, requestHeaders []communication.Header) (*communication.Response, error)

	// Post sends a POST request to the Ingenico ePayments platform and return the response.
	Post(resourceURI url.URL, requestHeaders []communication.Header, body string) (*communication.Response, error)

	// Put sends a PUT request to the Ingenico ePayments platform and return the response.
	Put(resourceURI url.URL, requestHeaders []communication.Header, body string) (*communication.Response, error)

	//IMPLEMENTATION logging.Capable INTERFACE

	// EnableLogging turns on logging using the given communicator logger.
	EnableLogging(communicatorLogger logging.CommunicatorLogger)

	// DisableLogging turns off logging.
	DisableLogging()

	//IMPLEMENTATION io.Closer INTERFACE

	// Close closes the connection of the Communicator
	Close() error
}
