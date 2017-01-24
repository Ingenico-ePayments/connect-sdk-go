package connectsdk

import (
	"reflect"
	"testing"

	"github.com/Ingenico-ePayments/connect-sdk-go/configuration"
	"github.com/Ingenico-ePayments/connect-sdk-go/defaultimpl"
)

var apiKeyID = "someKey"
var secretAPIKey = "someSecret"

func TestCreateConfiguration(t *testing.T) {
	defaultConfiguration := configuration.DefaultConfiguration(apiKeyID, secretAPIKey, "Test")

	configuration, err := CreateConfiguration(apiKeyID, secretAPIKey, "Test")
	if err != nil {
		t.Fatalf("TestCreateConfiguration: %v", err)
	}

	if configuration.APIEndpoint != defaultConfiguration.APIEndpoint {
		t.Fatalf("TestCreateConfiguration: APIEndpoint mismatch %v %v",
			configuration.APIEndpoint, defaultConfiguration.APIEndpoint)
	}

	if configuration.AuthorizationType != defaultConfiguration.AuthorizationType {
		t.Fatalf("TestCreateConfiguration: AuthorizationType mismatch %v %v",
			configuration.AuthorizationType, defaultConfiguration.AuthorizationType)
	}

	if configuration.IdleTimeout != defaultConfiguration.IdleTimeout {
		t.Fatalf("TestCreateConfiguration: IdleTimeout mismatch %v %v",
			configuration.IdleTimeout, defaultConfiguration.IdleTimeout)
	}

	if configuration.ConnectTimeout != defaultConfiguration.ConnectTimeout {
		t.Fatalf("TestCreateConfiguration: ConnectTimeout mismatch %v %v",
			configuration.ConnectTimeout, defaultConfiguration.ConnectTimeout)
	}

	if configuration.SocketTimeout != defaultConfiguration.SocketTimeout {
		t.Fatalf("TestCreateConfiguration: SocketTimeout mismatch %v %v",
			configuration.SocketTimeout, defaultConfiguration.SocketTimeout)
	}

	if configuration.MaxConnections != defaultConfiguration.MaxConnections {
		t.Fatalf("TestCreateConfiguration: MaxConnections mismatch %v %v",
			configuration.MaxConnections, defaultConfiguration.MaxConnections)
	}

	if configuration.APIKeyID != apiKeyID {
		t.Fatalf("TestCreateConfiguration: APIKeyID mismatch %v %v",
			configuration.APIKeyID, apiKeyID)
	}

	if configuration.SecretAPIKey != secretAPIKey {
		t.Fatalf("TestCreateConfiguration: SecretAPIKey mismatch %v %v",
			configuration.SecretAPIKey, secretAPIKey)
	}
}

func TestCreateCommunicator(t *testing.T) {
	marshaller, _ := defaultimpl.NewDefaultMarshaller()

	communicator, err := CreateCommunicator(apiKeyID, secretAPIKey, "Test")
	if err != nil {
		t.Fatalf("TestCreateCommunicator: %v", err)
	}
	if communicator.Marshaller() != marshaller {
		t.Fatalf("TestCreateCommunicator: marshaller mismatch %v %v",
			communicator.Marshaller(), marshaller)
	}

	session := communicator.Session()

	connection := session.Connection()
	if _, isDefaultConnection := connection.(*defaultimpl.DefaultConnection); !isDefaultConnection {
		t.Fatalf("TestCreateCommunicator: connection type mismatch %T", connection)
	}

	authenticator := session.Authenticator()
	if _, isDefaultAuthenticator := authenticator.(*defaultimpl.DefaultAuthenticator); !isDefaultAuthenticator {
		t.Fatalf("TestCreateCommunicator: authenticator type mismatch %T", authenticator)
	}

	authAuthType := reflect.Indirect(reflect.ValueOf(authenticator)).FieldByName("authType").String()
	if authAuthType != defaultimpl.V1HMAC {
		t.Fatalf("TestCreateCommunicator: auth type mismatch %v", authAuthType)
	}
	authAPIKeyID := reflect.Indirect(reflect.ValueOf(authenticator)).FieldByName("apiKeyID").String()
	if authAPIKeyID != apiKeyID {
		t.Fatalf("TestCreateCommunicator: apiKeyID mismatch %v", authAPIKeyID)
	}
	authSecretAPIKey := reflect.Indirect(reflect.ValueOf(authenticator)).FieldByName("secretAPIKey").String()
	if authSecretAPIKey != secretAPIKey {
		t.Fatalf("TestCreateCommunicator: secretAPIKey mismatch %v", authSecretAPIKey)
	}

	metaDataProvider := session.MetaDataProvider()
	headers := metaDataProvider.MetaDataHeaders()
	if len(headers) != 1 {
		t.Fatalf("TestCreateCommunicator: headers len mismatch %v", len(headers))
	}
	if headers[0].Name() != "X-GCS-ServerMetaInfo" {
		t.Fatalf("TestCreateCommunicator: header name mismatch %v", headers[0].Name())
	}
}
