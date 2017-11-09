// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

// RedirectDataBase represents class RedirectDataBase
type RedirectDataBase struct {
	RETURNMAC   *string `json:"RETURNMAC,omitempty"`
	RedirectURL *string `json:"redirectURL,omitempty"`
}

// NewRedirectDataBase constructs a new RedirectDataBase
func NewRedirectDataBase() *RedirectDataBase {
	return &RedirectDataBase{}
}
