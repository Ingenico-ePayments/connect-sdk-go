// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package product

// PaymentProductNetworksResponse represents class PaymentProductNetworksResponse
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentProductNetworksResponse
type PaymentProductNetworksResponse struct {
	Networks *[]string `json:"networks,omitempty"`
}

// NewPaymentProductNetworksResponse constructs a new PaymentProductNetworksResponse
func NewPaymentProductNetworksResponse() *PaymentProductNetworksResponse {
	return &PaymentProductNetworksResponse{}
}
