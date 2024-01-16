// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package product

// AuthenticationIndicator represents class AuthenticationIndicator
type AuthenticationIndicator struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewAuthenticationIndicator constructs a new AuthenticationIndicator
func NewAuthenticationIndicator() *AuthenticationIndicator {
	return &AuthenticationIndicator{}
}
