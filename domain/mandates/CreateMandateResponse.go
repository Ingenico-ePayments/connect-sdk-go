// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package mandates

// CreateMandateResponse represents class CreateMandateResponse
type CreateMandateResponse struct {
	Mandate        *MandateResponse       `json:"mandate,omitempty"`
	MerchantAction *MandateMerchantAction `json:"merchantAction,omitempty"`
}

// NewCreateMandateResponse constructs a new CreateMandateResponse
func NewCreateMandateResponse() *CreateMandateResponse {
	return &CreateMandateResponse{}
}
