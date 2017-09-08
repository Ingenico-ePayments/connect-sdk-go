package webhooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/Ingenico-ePayments/connect-sdk-go"
	"github.com/Ingenico-ePayments/connect-sdk-go/communicator"
	"github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/webhooks"
)

var (
	// ErrNilMarshaller occurs when the provided marshaller is nil
	ErrNilMarshaller = errors.New("nil marshaller")
	// ErrNilSecretKeyStore occurs when the provided secretKeyStore is nil
	ErrNilSecretKeyStore = errors.New("nil secretKeyStore")

	multipleHeaderFormat = `enocuntered multiple occurrences of header "%v"`
	headerNotFoundFormat = `could not find header "%v"`
	compareFailedFormat  = `failed to validate signature "%v"`
)

// Helper is the Ingenico ePayments platform webhooks helper. Immutable and thread-safe.
type Helper struct {
	marshaller     communicator.Marshaller
	secretKeyStore SecretKeyStore
}

// Unmarshal unmarshalls the given body and validates it using the requestHeaders
//
// Can return any of the following errors:
// SignatureValidationError if the body could not be validated succesfully
// APIVersionMismatchError if the resulting event has an API version that this version of the SDK does not support
func (h *Helper) Unmarshal(body string, requestHeaders []communication.Header) (*webhooks.Event, error) {
	err := h.validate(body, requestHeaders)
	if err != nil {
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

func (h *Helper) validate(body string, requestHeaders []communication.Header) error {
	signature, err := getHeaderValue(requestHeaders, "X-GCS-Signature")
	if err != nil {
		return err
	}

	keyID, err := getHeaderValue(requestHeaders, "X-GCS-KeyId")
	if err != nil {
		return err
	}

	secretKey, err := h.secretKeyStore.GetSecretKey(keyID)
	if err != nil {
		return err
	}

	mac := hmac.New(sha256.New, []byte(secretKey))

	_, err = mac.Write([]byte(body))
	if err != nil {
		return err
	}

	unencodedResult := mac.Sum(nil)

	expectedSignature := base64.StdEncoding.EncodeToString(unencodedResult)

	isValid := areEqualSignatures(signature, expectedSignature)

	if !isValid {
		sve, err := NewSignatureValidationError(fmt.Sprintf(compareFailedFormat, signature))
		if err != nil {
			return err
		}

		return sve
	}

	return nil
}

// NewHelper creates an Helper with the given marshaller and secretKeyStore
func NewHelper(marshaller communicator.Marshaller, secretKeyStore SecretKeyStore) (*Helper, error) {
	if marshaller == nil {
		return nil, ErrNilMarshaller
	}

	if secretKeyStore == nil {
		return nil, ErrNilSecretKeyStore
	}

	return &Helper{marshaller, secretKeyStore}, nil
}

func areEqualSignatures(signature, expectedSignature string) bool {
	signatureRunes := []rune(signature)
	expectedSignatureRunes := []rune(expectedSignature)

	length := len(signatureRunes)
	expectedLength := len(expectedSignatureRunes)

	limit := max(256, max(length, expectedLength))

	result := true

	for i := 0; i < limit; i++ {
		if i < length && i < expectedLength {
			result = result && (signatureRunes[i] == expectedSignatureRunes[i])
		} else {
			if i >= length && i >= expectedLength {
				result = result && true
			} else {
				result = result && false
			}
		}
	}

	return result
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

func getHeaderValue(headers []communication.Header, name string) (string, error) {
	lowerName := strings.ToLower(name)
	var value string

	found := false

	for _, header := range headers {
		if strings.ToLower(header.Name()) == lowerName {
			if found {
				sve, err := NewSignatureValidationError(fmt.Sprintf(multipleHeaderFormat, name))
				if err != nil {
					return "", err
				}

				return "", sve
			}

			found, value = true, header.Value()
		}
	}

	if !found {
		sve, err := NewSignatureValidationError(fmt.Sprintf(headerNotFoundFormat, name))
		if err != nil {
			return "", err
		}

		return "", sve
	}

	return value, nil
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
