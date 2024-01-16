// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// CreateResponse represents class CreatePaymentResponse
type CreateResponse struct {
	CreationOutput *CreationOutput `json:"creationOutput,omitempty"`
	MerchantAction *MerchantAction `json:"merchantAction,omitempty"`
	Payment        *Payment        `json:"payment,omitempty"`
}

// NewCreateResponse constructs a new CreateResponse
func NewCreateResponse() *CreateResponse {
	return &CreateResponse{}
}
