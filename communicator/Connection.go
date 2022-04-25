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

	// CloseIdleConnections closes all HTTP connections that have been idle for the specified time. This should also include
	// all expired HTTP connections.
	// timespan represents the time spent idle
	CloseIdleConnections(time time.Duration)

	// CloseExpiredConnections closes all expired HTTP connections.
	CloseExpiredConnections()

	//IMPLEMENTATION Connection INTERFACE

	// Get sends a GET request to the Ingenico ePayments platform and calls the given response handler with the response.
	Get(resourceURI url.URL, requestHeaders []communication.Header, respHandler communication.ResponseHandler) (interface{}, error)

	// Delete sends a DELETE request to the Ingenico ePayments platform and calls the given response handler with the response.
	Delete(resourceURI url.URL, requestHeaders []communication.Header, respHandler communication.ResponseHandler) (interface{}, error)

	// Post sends a POST request to the Ingenico ePayments platform and calls the given response handler with the response.
	Post(resourceURI url.URL, requestHeaders []communication.Header, body string, respHandler communication.ResponseHandler) (interface{}, error)

	// PostMultipart sends a multipart/form-data POST request to the Ingenico ePayments platform and calls the given response handler with the response.
	PostMultipart(resourceURI url.URL, requestHeaders []communication.Header, body *communication.MultipartFormDataObject, respHandler communication.ResponseHandler) (interface{}, error)

	// Put sends a PUT request to the Ingenico ePayments platform and calls the given response handler with the response.
	Put(resourceURI url.URL, requestHeaders []communication.Header, body string, respHandler communication.ResponseHandler) (interface{}, error)

	// PutMultipart sends a multipart/form-data PUT request to the Ingenico ePayments platform and calls the given response handler with the response.
	PutMultipart(resourceURI url.URL, requestHeaders []communication.Header, body *communication.MultipartFormDataObject, respHandler communication.ResponseHandler) (interface{}, error)

	//IMPLEMENTATION logging.Capable INTERFACE

	// EnableLogging turns on logging using the given communicator logger.
	EnableLogging(communicatorLogger logging.CommunicatorLogger)

	// DisableLogging turns off logging.
	DisableLogging()

	//IMPLEMENTATION io.Closer INTERFACE

	// Close closes the connection of the Communicator
	Close() error
}
