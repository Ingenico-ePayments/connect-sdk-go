// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// CreateResponse represents class CreatePaymentResponse
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_CreatePaymentResponse
type CreateResponse struct {
	CreationOutput *CreationOutput `json:"creationOutput,omitempty"`
	MerchantAction *MerchantAction `json:"merchantAction,omitempty"`
	Payment        *Payment        `json:"payment,omitempty"`
}

// NewCreateResponse constructs a new CreateResponse
func NewCreateResponse() *CreateResponse {
	return &CreateResponse{}
}
