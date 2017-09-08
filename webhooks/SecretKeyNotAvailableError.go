package webhooks

import "fmt"

var (
	keyNotFoundErrorFormat = `could not find secret key for key id "%v"`
)

// SecretKeyNotAvailableError represents an error that causes a secret key to not be available.
type SecretKeyNotAvailableError struct {
	keyID string
	message string
}

// Error implements the Error interface
func (ske *SecretKeyNotAvailableError) Error() string {
	return ske.message
}

// KeyID returns the keyID which produced the error
func (ske *SecretKeyNotAvailableError) KeyID() string {
	return ske.keyID
}

// NewSecretKeyNotAvailableError creates a SecretKeyNotAvailableError with the given keyID
func NewSecretKeyNotAvailableError(keyID string) (*SecretKeyNotAvailableError, error) {
	message := fmt.Sprintf(keyNotFoundErrorFormat, keyID)

	return &SecretKeyNotAvailableError{keyID, message}, nil
}
