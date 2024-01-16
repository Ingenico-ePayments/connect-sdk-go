// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package definitions

// KeyValuePair represents class KeyValuePair
type KeyValuePair struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewKeyValuePair constructs a new KeyValuePair
func NewKeyValuePair() *KeyValuePair {
	return &KeyValuePair{}
}
