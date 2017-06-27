// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package publickey

// PublicKey represents class PublicKey
type PublicKey struct {
	KeyID     *string `json:"keyId,omitempty"`
	PublicKey *string `json:"publicKey,omitempty"`
}

// NewPublicKey constructs a new PublicKey
func NewPublicKey() *PublicKey {
	return &PublicKey{}
}
