package validation

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"
)

var (
	// ErrNilSecretKeyStore occurs when the provided secretKeyStore is nil
	ErrNilSecretKeyStore = errors.New("nil secretKeyStore")

	multipleHeaderFormat = `encountered multiple occurrences of header "%v"`
	headerNotFoundFormat = `could not find header "%v"`
	compareFailedFormat  = `failed to validate signature "%v"`
)

// SignatureValidator is a validator for webhooks signatures. Immutable and thread-safe.
type SignatureValidator struct {
	secretKeyStore SecretKeyStore
}

// SecretKeyStore returns the configured secret key store
func (v *SignatureValidator) SecretKeyStore() SecretKeyStore {
	return v.secretKeyStore
}

// Validate validates the given body using the given request headers
func (v *SignatureValidator) Validate(body string, requestHeaders []communication.Header) error {
	signature, err := getHeaderValue(requestHeaders, "X-GCS-Signature")
	if err != nil {
		return err
	}

	keyID, err := getHeaderValue(requestHeaders, "X-GCS-KeyId")
	if err != nil {
		return err
	}

	secretKey, err := v.secretKeyStore.GetSecretKey(keyID)
	if err != nil {
		return err
	}

	mac := hmac.New(sha256.New, []byte(secretKey))

	_, err = mac.Write([]byte(body))
	if err != nil {
		return err
	}

	unencodedResult := mac.Sum(nil)

	encoder := base64.StdEncoding
	expectedSignature := make([]byte, encoder.EncodedLen(len(unencodedResult)))
	encoder.Encode(expectedSignature, unencodedResult)

	isValid := subtle.ConstantTimeCompare([]byte(signature), expectedSignature) == 1
	if !isValid {
		return NewSignatureValidationError(fmt.Sprintf(compareFailedFormat, signature))
	}

	return nil
}

// NewSignatureValidator creates a signature validator with the given secretKeyStore
func NewSignatureValidator(secretKeyStore SecretKeyStore) (*SignatureValidator, error) {
	if secretKeyStore == nil {
		return nil, ErrNilSecretKeyStore
	}

	return &SignatureValidator{secretKeyStore}, nil
}

func getHeaderValue(headers []communication.Header, name string) (string, error) {
	lowerName := strings.ToLower(name)
	var value string

	found := false

	for _, header := range headers {
		if strings.ToLower(header.Name()) == lowerName {
			if found {
				sve := NewSignatureValidationError(fmt.Sprintf(multipleHeaderFormat, name))

				return "", sve
			}

			found, value = true, header.Value()
		}
	}

	if !found {
		sve := NewSignatureValidationError(fmt.Sprintf(headerNotFoundFormat, name))

		return "", sve
	}

	return value, nil
}
