package webhooks

import "github.com/Ingenico-ePayments/connect-sdk-go/defaultimpl"

// CreateHelperBuilder creates a HelperBuilder that will use the given secretKeyStore
func CreateHelperBuilder(secretKeyStore SecretKeyStore) (*HelperBuilder, error) {
	marshaller, err := defaultimpl.NewDefaultMarshaller()
	if err != nil {
		return nil, err
	}

	return NewHelperBuilder().WithMarshaller(marshaller).WithSecretKeyStore(secretKeyStore), nil
}

// CreateHelper creates a Helper that will use the given secretKeyStore
func CreateHelper(secretKeyStore SecretKeyStore) (*Helper, error) {
	helperBuilder, err := CreateHelperBuilder(secretKeyStore);
	if err != nil {
		return nil, err
	}

	return helperBuilder.Build()
}
