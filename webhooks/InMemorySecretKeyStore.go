package webhooks

import (
	"errors"
	"strings"
	"sync"
)

var (
	// ErrInvalidKey occurs when the keyID is empty
	ErrInvalidKey = errors.New("invalid keyID")
	// ErrInvalidSecret occurs when the given secretKey is empty
	ErrInvalidSecret = errors.New("invalid secret")
)

// InMemorySecretKeyStore is an in-memory secret key store.
// This implementation can be used in applications where secret keys can be specified at application startup.
// Thread-safe.
type InMemorySecretKeyStore struct{}

var keyStoreMap = map[string]string{}
var keyStoreLock = sync.RWMutex{}

// GetSecretKey returns the secretKey associated with the given keyID
//
// Can return any of the following errors:
// SecretKeyNotAvailableError if there is no secretKey with the given keyID
func (ks *InMemorySecretKeyStore) GetSecretKey(keyID string) (string, error) {
	keyStoreLock.RLock()
	defer keyStoreLock.RUnlock()

	value, exists := keyStoreMap[keyID]

	if !exists {
		ske, err := NewSecretKeyNotAvailableError(keyID)
		if err != nil {
			return "", err
		}

		return "", ske
	}

	return value, nil
}

// StoreSecretKey stores the given keyID, secretKey pair
//
// Can return any of the following errors:
// ErrInvalidKey if the keyID is empty
// ErrInvalidSecret if the secretKey is empty
func (ks *InMemorySecretKeyStore) StoreSecretKey(keyID, secretKey string) error {
	if strings.TrimSpace(keyID) == "" {
		return ErrInvalidKey
	}
	if strings.TrimSpace(secretKey) == "" {
		return ErrInvalidSecret
	}

	keyStoreLock.Lock()
	defer keyStoreLock.Unlock()

	keyStoreMap[keyID] = secretKey

	return nil
}

// RemoveSecretKey removes the given keyID and it's associated secretKey
func (ks *InMemorySecretKeyStore) RemoveSecretKey(keyID string) {
	keyStoreLock.Lock()
	defer keyStoreLock.Unlock()

	delete(keyStoreMap, keyID)
}

// Clear empties the key store
func (ks *InMemorySecretKeyStore) Clear() {
	keyStoreLock.Lock()
	defer keyStoreLock.Unlock()

	keyStoreMap = map[string]string{}
}

// NewInMemorySecretKeyStore returns the singleton instance of the InMemorySecretKeyStore
func NewInMemorySecretKeyStore() (*InMemorySecretKeyStore, error) {
	return &InMemorySecretKeyStore{}, nil
}
