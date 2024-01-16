package webhooks

import (
	"errors"
	"github.com/Ingenico-ePayments/connect-sdk-go"
	"github.com/Ingenico-ePayments/connect-sdk-go/communicator"
	"github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/webhooks"
	"github.com/Ingenico-ePayments/connect-sdk-go/webhooks/validation"
)

var (
	// ErrNilMarshaller occurs when the provided marshaller is nil
	ErrNilMarshaller = errors.New("nil marshaller")
	// ErrNilSecretKeyStore occurs when the provided secretKeyStore is nil
	// Deprecated: use validation.ErrNilSecretKeyStore instead
	ErrNilSecretKeyStore = validation.ErrNilSecretKeyStore
)

// Helper is the Ingenico ePayments platform webhooks helper. Immutable and thread-safe.
type Helper struct {
	marshaller         communicator.Marshaller
	signatureValidator *validation.SignatureValidator
}

// Unmarshal unmarshalls the given body and validates it using the requestHeaders
//
// Can return any of the following errors:
// SignatureValidationError if the body could not be validated successfully
// APIVersionMismatchError if the resulting event has an API version that this version of the SDK does not support
func (h *Helper) Unmarshal(body string, requestHeaders []communication.Header) (*webhooks.Event, error) {
	err := h.signatureValidator.Validate(body, requestHeaders)
	if err != nil {
		if _, ok := err.(*validation.SignatureValidationError); ok {
			// Return the right type but with the same error message
			sve, _ := NewSignatureValidationError(err.Error())
			return nil, sve
		}

		return nil, err
	}

	event, err := webhooks.NewEvent()
	if err != nil {
		return nil, err
	}

	err = h.marshaller.Unmarshal(body, event)
	if err != nil {
		return nil, err
	}

	err = validateAPIVersion(event)
	if err != nil {
		return nil, err
	}

	return event, nil
}

// NewHelper creates an Helper with the given marshaller and secretKeyStore
func NewHelper(marshaller communicator.Marshaller, secretKeyStore SecretKeyStore) (*Helper, error) {
	if marshaller == nil {
		return nil, ErrNilMarshaller
	}

	signatureValidator, err := validation.NewSignatureValidator(secretKeyStore)
	if err != nil {
		return nil, err
	}

	return &Helper{marshaller, signatureValidator}, nil
}

func validateAPIVersion(event *webhooks.Event) error {
	if *event.APIVersion != connectsdk.APIVersion {
		ame, err := NewAPIVersionMismatchError(*event.APIVersion, connectsdk.APIVersion)
		if err != nil {
			return err
		}

		return ame
	}

	return nil
}
