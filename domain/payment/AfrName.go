// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// AfrName represents class AfrName
type AfrName struct {
	FirstName *string `json:"firstName,omitempty"`
	Surname   *string `json:"surname,omitempty"`
}

// NewAfrName constructs a new AfrName
func NewAfrName() *AfrName {
	return &AfrName{}
}
