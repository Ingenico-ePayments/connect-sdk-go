// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// CustomerAccountAuthentication represents class CustomerAccountAuthentication
type CustomerAccountAuthentication struct {
	Data         *string `json:"data,omitempty"`
	Method       *string `json:"method,omitempty"`
	UtcTimestamp *string `json:"utcTimestamp,omitempty"`
}

// NewCustomerAccountAuthentication constructs a new CustomerAccountAuthentication
func NewCustomerAccountAuthentication() *CustomerAccountAuthentication {
	return &CustomerAccountAuthentication{}
}
