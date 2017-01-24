package connectsdk

import (
	"net/url"
	"testing"
	"encoding/base64"

	"github.com/Ingenico-ePayments/connect-sdk-go/configuration"
	"github.com/Ingenico-ePayments/connect-sdk-go/defaultimpl"
)

var testConfiguration = configuration.CommunicatorConfiguration{
	APIEndpoint: url.URL{
		Scheme: "https",
		Host:   "api-sandbox.globalcollect.com",
	},
	AuthorizationType: defaultimpl.V1HMAC,
	MaxConnections:    100,
	APIKeyID:          "someKey",
	SecretAPIKey:      "someSecret",
}

func TestWithClientMetaInfo(t *testing.T) {
	conf := testConfiguration
	client1, err := CreateClientFromConfiguration(&conf)
	if err != nil {
		t.Fatalf("TestWithClientMetaInfo: %v", err)
	}
	if client1.apiResource == nil {
		t.Fatal("TestWithClientMetaInfo: nil apiResource")
	}

	headers := client1.apiResource.ClientHeaders()
	if headers != nil {
		t.Fatal("TestWithClientMetaInfo: headers not nil")
	}

	client2, err := client1.WithClientMetaInfo("")
	if err != nil {
		t.Fatalf("TestWithClientMetaInfo: %v", err)
	}
	if client1 != client2 {
		t.Fatal("TestWithClientMetaInfo: clients not equal")
	}

	marshaller, _ := defaultimpl.NewDefaultMarshaller()
	clientMetaInfo, err := marshaller.Marshal(map[string]string{
		"test": "test",
	})
	if err != nil {
		t.Fatalf("TestWithClientMetaInfo: %v", err)
	}
	client3, err := client1.WithClientMetaInfo(clientMetaInfo)
	if err != nil {
		t.Fatalf("TestWithClientMetaInfo: %v", err)
	}
	if client3 == client1 {
		t.Fatal("TestWithClientMetaInfo: clients equal")
	}
	headers = client3.apiResource.ClientHeaders()
	if len(headers) != 1 {
		t.Fatalf("TestWithClientMetaInfo: invalid headers len %d", len(headers))
	}
	headerValue := base64.StdEncoding.EncodeToString([]byte(clientMetaInfo))
	if headers[0].Value() != headerValue {
		t.Fatalf("TestWithClientMetaInfo: header value mismatch '%v' '%v'", headers[0].Value(), headerValue)
	}

	client4, err := client3.WithClientMetaInfo(clientMetaInfo)
	if err != nil {
		t.Fatalf("TestWithClientMetaInfo: %v", err)
	}
	if client3 != client4 {
		t.Fatal("TestWithClientMetaInfo: clients not equal")
	}

	client5, err := client3.WithClientMetaInfo("")
	if err != nil {
		t.Fatal("TestWithClientMetaInfo: clients not equal")
	}
	if client3 == client5 {
		t.Fatal("TestWithClientMetaInfo: clients equal")
	}
}
