package webhooks

// SecretKeyStore represents a store of secret keys.
// Implementations could store secret keys in a database, on disk, etc.
// Thread-safe.
type SecretKeyStore interface {
	// GetSecretKey returns the secretKey for the given keyID
	GetSecretKey(string) (string, error)
}
