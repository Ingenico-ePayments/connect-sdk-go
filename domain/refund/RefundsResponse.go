// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package refund

// RefundsResponse represents class RefundsResponse
type RefundsResponse struct {
	Refunds *[]Result `json:"refunds,omitempty"`
}

// NewRefundsResponse constructs a new RefundsResponse
func NewRefundsResponse() *RefundsResponse {
	return &RefundsResponse{}
}
