package webhooks

import (
	"github.com/Ingenico-ePayments/connect-sdk-go/defaultimpl"
	"testing"
)

func TestCreateHelper(t *testing.T) {
	store, err := NewInMemorySecretKeyStore()
	if err != nil {
		t.Fatal(err)
	}

	helper, err := CreateHelper(store)
	if err != nil {
		t.Fatal(err)
	}

	marshaller, err := defaultimpl.NewDefaultMarshaller()
	if err != nil {
		t.Fatal(err)
	}

	if helper.marshaller != marshaller {
		t.Fatalf("marshaller mismatch %v %v", helper.marshaller, marshaller)
	}
	if helper.signatureValidator.SecretKeyStore() != store {
		t.Fatalf("secretKeyStore mismatch %v %v", helper.signatureValidator.SecretKeyStore(), store)
	}
}
