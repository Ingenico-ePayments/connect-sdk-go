package communicator

import (
	"net/url"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"
)

// Authenticator is the interface used to sign requests to the Ingenico ePayments platform. Thread-safe.
type Authenticator interface {

	// CreateSimpleAuthenticationSignature creates a signature for the simple security model.
	// It returns the simple authentication signature.
	// Note that the list of Request headers may not be modified and may not contain headers with
	// the same name.
	CreateSimpleAuthenticationSignature(httpMethod string, resourceURI url.URL, requestHeaders []communication.Header) (string, error)
}
