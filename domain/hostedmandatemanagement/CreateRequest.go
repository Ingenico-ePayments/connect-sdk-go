// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package hostedmandatemanagement

// CreateRequest represents class CreateHostedMandateManagementRequest
type CreateRequest struct {
	CreateMandateInfo                    *HostedMandateInfo `json:"createMandateInfo,omitempty"`
	HostedMandateManagementSpecificInput *SpecificInput     `json:"hostedMandateManagementSpecificInput,omitempty"`
}

// NewCreateRequest constructs a new CreateRequest
func NewCreateRequest() *CreateRequest {
	return &CreateRequest{}
}
