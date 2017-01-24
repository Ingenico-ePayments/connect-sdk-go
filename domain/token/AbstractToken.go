// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package token

// AbstractToken represents class AbstractToken
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_AbstractToken
type AbstractToken struct {
	Alias *string `json:"alias,omitempty"`
}

// NewAbstractToken constructs a new AbstractToken
func NewAbstractToken() *AbstractToken {
	return &AbstractToken{}
}
