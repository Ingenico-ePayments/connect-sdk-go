package webhooks

import "github.com/Ingenico-ePayments/connect-sdk-go/webhooks/validation"

// SecretKeyStore represents a store of secret keys.
// Implementations could store secret keys in a database, on disk, etc.
// Thread-safe.
type SecretKeyStore validation.SecretKeyStore
