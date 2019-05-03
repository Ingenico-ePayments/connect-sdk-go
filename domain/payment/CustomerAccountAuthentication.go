// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// CustomerAccountAuthentication represents class CustomerAccountAuthentication
type CustomerAccountAuthentication struct {
	Method       *string `json:"method,omitempty"`
	UtcTimestamp *string `json:"utcTimestamp,omitempty"`
}

// NewCustomerAccountAuthentication constructs a new CustomerAccountAuthentication
func NewCustomerAccountAuthentication() *CustomerAccountAuthentication {
	return &CustomerAccountAuthentication{}
}
