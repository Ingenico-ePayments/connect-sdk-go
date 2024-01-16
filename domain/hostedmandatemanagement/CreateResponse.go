// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package hostedmandatemanagement

// CreateResponse represents class CreateHostedMandateManagementResponse
type CreateResponse struct {
	RETURNMAC                 *string `json:"RETURNMAC,omitempty"`
	HostedMandateManagementID *string `json:"hostedMandateManagementId,omitempty"`
	PartialRedirectURL        *string `json:"partialRedirectUrl,omitempty"`
}

// NewCreateResponse constructs a new CreateResponse
func NewCreateResponse() *CreateResponse {
	return &CreateResponse{}
}
