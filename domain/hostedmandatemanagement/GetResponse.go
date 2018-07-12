// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package hostedmandatemanagement

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/mandates"

// GetResponse represents class GetHostedMandateManagementResponse
type GetResponse struct {
	Mandate *mandates.MandateResponse `json:"mandate,omitempty"`
	Status  *string                   `json:"status,omitempty"`
}

// NewGetResponse constructs a new GetResponse
func NewGetResponse() *GetResponse {
	return &GetResponse{}
}
