// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package hostedcheckout

// GetResponse represents class GetHostedCheckoutResponse
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_GetHostedCheckoutResponse
type GetResponse struct {
	CreatedPaymentOutput *CreatedPaymentOutput `json:"createdPaymentOutput,omitempty"`
	Status               *string               `json:"status,omitempty"`
}

// NewGetResponse constructs a new GetResponse
func NewGetResponse() *GetResponse {
	return &GetResponse{}
}
