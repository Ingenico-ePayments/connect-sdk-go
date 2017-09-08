package webhooks

import "github.com/Ingenico-ePayments/connect-sdk-go/communicator"

// HelperBuilder is used to build Helper objects
type HelperBuilder struct {
	Marshaller     communicator.Marshaller
	SecretKeyStore SecretKeyStore
}

// WithMarshaller sets the marshaller
func (h *HelperBuilder) WithMarshaller(marshaller communicator.Marshaller) *HelperBuilder {
	h.Marshaller = marshaller

	return h
}

// WithSecretKeyStore sets the secretKeyStore
func (h *HelperBuilder) WithSecretKeyStore(secretKeyStore SecretKeyStore) *HelperBuilder {
	h.SecretKeyStore = secretKeyStore

	return h
}

// Build creates the Helper object
func (h *HelperBuilder) Build() (*Helper, error) {
	return NewHelper(h.Marshaller, h.SecretKeyStore)
}

// NewHelperBuilder creates a HelperBuilder object
func NewHelperBuilder() *HelperBuilder {
	return &HelperBuilder{}
}
