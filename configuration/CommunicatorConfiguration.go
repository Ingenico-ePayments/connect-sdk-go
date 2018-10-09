package configuration

import (
	"net/url"
	"time"

	"github.com/Ingenico-ePayments/connect-sdk-go/defaultimpl"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/metadata"
)

// CommunicatorConfiguration represents the configuration to be used by a Communicator
type CommunicatorConfiguration struct {
	// APIEndpoint represents the API endpoint of for the communicator
	// See https://epayments-api.developer-ingenico.com/s2sapi/v1/en_US/go/endpoints.html
	APIEndpoint url.URL
	// ConnectTimeout represents the connect timeout
	ConnectTimeout time.Duration
	// SocketTimeout represents the request timeout
	SocketTimeout time.Duration
	// IdleTimeout represents the idle connection timeout
	IdleTimeout time.Duration
	// KeepAliveTimeout represents the HTTP KeepAlive interval
	KeepAliveTimeout time.Duration
	// MaxConnections represents the maximum amount of concurrent pooled connections
	MaxConnections int
	// AuthorizationType represents authorizationType used to sign the requests
	AuthorizationType defaultimpl.AuthorizationType
	// APIKeyID represents an identifier for the secret API key
	APIKeyID string
	// SecretAPIKey represents a shared secret
	SecretAPIKey string
	// Proxy represents the URL for the connection proxy
	Proxy *url.URL
	// Integrator represents the integrator name
	Integrator string
	// ShoppingCartExtension represents the shopping cart extension used in the MetaData headers
	ShoppingCartExtension *metadata.ShoppingCartExtension
}

// The default configuration used by the factory is the following:
// APIEndpoint: world.api-ingenico.com
// ConnectTimeout: 5 seconds
// SocketTimeout: 30 seconds
// IdleTimeout: 5 seconds
// KeepAliveTimeout: 30 seconds
// MaxConnections: 10
// AuthorizationType: V1HMAC
// Proxy: none
var defaultConfiguration = CommunicatorConfiguration{
	APIEndpoint: url.URL{
		Scheme: "https",
		Host:   "world.api-ingenico.com",
	},
	ConnectTimeout:    5 * time.Second,
	SocketTimeout:     30 * time.Second,
	IdleTimeout:       5 * time.Second,
	KeepAliveTimeout:  30 * time.Second,
	MaxConnections:    10,
	AuthorizationType: defaultimpl.V1HMAC,
}

// DefaultConfiguration returns the default communicator configuration
func DefaultConfiguration(apiKeyID, secretAPIKey, integrator string) CommunicatorConfiguration {
	customizedConfiguration := defaultConfiguration

	customizedConfiguration.Integrator = integrator
	customizedConfiguration.SecretAPIKey = secretAPIKey
	customizedConfiguration.APIKeyID = apiKeyID

	return customizedConfiguration
}
