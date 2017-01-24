// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package hostedcheckout

// CreateResponse represents class CreateHostedCheckoutResponse
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_CreateHostedCheckoutResponse
type CreateResponse struct {
	RETURNMAC          *string   `json:"RETURNMAC,omitempty"`
	HostedCheckoutID   *string   `json:"hostedCheckoutId,omitempty"`
	InvalidTokens      *[]string `json:"invalidTokens,omitempty"`
	PartialRedirectURL *string   `json:"partialRedirectUrl,omitempty"`
}

// NewCreateResponse constructs a new CreateResponse
func NewCreateResponse() *CreateResponse {
	return &CreateResponse{}
}
