package connectsdk

import (
	"github.com/Ingenico-ePayments/connect-sdk-go/communicator"
	"github.com/Ingenico-ePayments/connect-sdk-go/configuration"
	"github.com/Ingenico-ePayments/connect-sdk-go/defaultimpl"
)

// CreateConfiguration creates a CommunicatorConfiguration with default settings and the given apiKeyID and secretAPIKey
func CreateConfiguration(apiKeyID, secretAPIKey, integrator string) (*configuration.CommunicatorConfiguration, error) {
	newConf := configuration.DefaultConfiguration(apiKeyID, secretAPIKey, integrator)

	return &newConf, nil
}

// CreateSessionBuilder creates a SessionBuilder with default settings and the given apiKeyID and secretAPIKey
func CreateSessionBuilder(apiKeyID, secretAPIKey, integrator string) (*communicator.SessionBuilder, error) {
	configuration, err := CreateConfiguration(apiKeyID, secretAPIKey, integrator)
	if err != nil {
		return nil, err
	}

	return CreateSessionBuilderFromConfiguration(configuration)
}

// CreateSessionBuilderFromConfiguration creates a SessionBuilder with the given CommunicatorConfiguration
func CreateSessionBuilderFromConfiguration(configuration *configuration.CommunicatorConfiguration) (*communicator.SessionBuilder, error) {
	builder, err := communicator.NewSessionBuilder()
	if err != nil {
		return nil, err
	}

	connection, err := defaultimpl.NewDefaultConnection(configuration.SocketTimeout,
		configuration.ConnectTimeout,
		configuration.KeepAliveTimeout,
		configuration.IdleTimeout,
		configuration.MaxConnections,
		configuration.Proxy)
	if err != nil {
		return nil, err
	}

	metaDataProviderBuilder := communicator.NewMetaDataProviderBuilder(configuration.Integrator)
	metaDataProviderBuilder.ShoppingCartExtension = configuration.ShoppingCartExtension

	metaDataProvider, err := metaDataProviderBuilder.Build()
	if err != nil {
		return nil, err
	}

	authenticator, err := defaultimpl.NewDefaultAuthenticator(configuration.AuthorizationType,
		configuration.APIKeyID,
		configuration.SecretAPIKey)
	if err != nil {
		return nil, err
	}

	return builder.WithAPIEndpoint(&configuration.APIEndpoint).WithConnection(connection).WithMetaDataProvider(metaDataProvider).WithAuthenticator(authenticator), nil
}

// CreateCommunicator creates a Communicator with default settings and the given apiKeyID and secretAPIKey
func CreateCommunicator(apiKeyID, secretAPIKey, integrator string) (*communicator.Communicator, error) {
	configuration, err := CreateConfiguration(apiKeyID, secretAPIKey, integrator)
	if err != nil {
		return nil, err
	}

	return CreateCommunicatorFromConfiguration(configuration)
}

// CreateCommunicatorFromConfiguration creates a Communicator with the given CommunicatorConfiguration
func CreateCommunicatorFromConfiguration(configuration *configuration.CommunicatorConfiguration) (*communicator.Communicator, error) {
	builder, err := CreateSessionBuilderFromConfiguration(configuration)
	if err != nil {
		return nil, err
	}

	session, err := builder.Build()
	if err != nil {
		return nil, err
	}

	return CreateCommunicatorFromSession(session)
}

// CreateCommunicatorFromSession creates a Communicator with the given Session
func CreateCommunicatorFromSession(session *communicator.Session) (*communicator.Communicator, error) {
	marshaller, err := defaultimpl.NewDefaultMarshaller()
	if err != nil {
		return nil, err
	}

	return communicator.NewCommunicator(session, marshaller)
}

// CreateClient creates a Client with default settings and the given apiKeyID and secretAPIKey
func CreateClient(apiKeyID, secretAPIKey, integrator string) (*Client, error) {
	communicator, err := CreateCommunicator(apiKeyID, secretAPIKey, integrator)
	if err != nil {
		return nil, err
	}

	return CreateClientFromCommunicator(communicator)
}

// CreateClientFromConfiguration creates a Client with the given CommunicatorConfiguration
func CreateClientFromConfiguration(configuration *configuration.CommunicatorConfiguration) (*Client, error) {
	communicator, err := CreateCommunicatorFromConfiguration(configuration)
	if err != nil {
		return nil, err
	}

	return CreateClientFromCommunicator(communicator)
}

// CreateClientFromSession creates a Client with the given Session
func CreateClientFromSession(session *communicator.Session) (*Client, error) {
	communicator, err := CreateCommunicatorFromSession(session)
	if err != nil {
		return nil, err
	}

	return CreateClientFromCommunicator(communicator)
}

// CreateClientFromCommunicator creates a Client with the given Communicator
func CreateClientFromCommunicator(communicator *communicator.Communicator) (*Client, error) {
	return NewClient(communicator)
}
