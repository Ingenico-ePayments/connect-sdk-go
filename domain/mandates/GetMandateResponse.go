// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package mandates

// GetMandateResponse represents class GetMandateResponse
type GetMandateResponse struct {
	Mandate *MandateResponse `json:"mandate,omitempty"`
}

// NewGetMandateResponse constructs a new GetMandateResponse
func NewGetMandateResponse() *GetMandateResponse {
	return &GetMandateResponse{}
}
