// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package publickey

// PublicKey represents class PublicKey
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PublicKey
type PublicKey struct {
	KeyID     *string `json:"keyId,omitempty"`
	PublicKey *string `json:"publicKey,omitempty"`
}

// NewPublicKey constructs a new PublicKey
func NewPublicKey() *PublicKey {
	return &PublicKey{}
}
