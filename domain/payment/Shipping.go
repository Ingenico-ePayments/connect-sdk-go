// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// Shipping represents class Shipping
type Shipping struct {
	EmailAddress *string `json:"emailAddress,omitempty"`
}

// NewShipping constructs a new Shipping
func NewShipping() *Shipping {
	return &Shipping{}
}
