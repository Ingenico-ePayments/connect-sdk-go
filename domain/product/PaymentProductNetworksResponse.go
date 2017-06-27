// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package product

// PaymentProductNetworksResponse represents class PaymentProductNetworksResponse
type PaymentProductNetworksResponse struct {
	Networks *[]string `json:"networks,omitempty"`
}

// NewPaymentProductNetworksResponse constructs a new PaymentProductNetworksResponse
func NewPaymentProductNetworksResponse() *PaymentProductNetworksResponse {
	return &PaymentProductNetworksResponse{}
}
