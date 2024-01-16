// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package hostedcheckout

// GetResponse represents class GetHostedCheckoutResponse
type GetResponse struct {
	CreatedPaymentOutput *CreatedPaymentOutput `json:"createdPaymentOutput,omitempty"`
	Status               *string               `json:"status,omitempty"`
}

// NewGetResponse constructs a new GetResponse
func NewGetResponse() *GetResponse {
	return &GetResponse{}
}
