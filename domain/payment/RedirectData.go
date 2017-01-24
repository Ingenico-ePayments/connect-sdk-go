// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// RedirectData represents class RedirectData
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_RedirectData
type RedirectData struct {
	RETURNMAC   *string `json:"RETURNMAC,omitempty"`
	RedirectURL *string `json:"redirectURL,omitempty"`
}

// NewRedirectData constructs a new RedirectData
func NewRedirectData() *RedirectData {
	return &RedirectData{}
}
