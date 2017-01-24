// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package definitions

// KeyValuePair represents class KeyValuePair
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_KeyValuePair
type KeyValuePair struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewKeyValuePair constructs a new KeyValuePair
func NewKeyValuePair() *KeyValuePair {
	return &KeyValuePair{}
}
