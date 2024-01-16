// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package hostedcheckout

// CreateResponse represents class CreateHostedCheckoutResponse
type CreateResponse struct {
	RETURNMAC          *string   `json:"RETURNMAC,omitempty"`
	HostedCheckoutID   *string   `json:"hostedCheckoutId,omitempty"`
	InvalidTokens      *[]string `json:"invalidTokens,omitempty"`
	MerchantReference  *string   `json:"merchantReference,omitempty"`
	PartialRedirectURL *string   `json:"partialRedirectUrl,omitempty"`
}

// NewCreateResponse constructs a new CreateResponse
func NewCreateResponse() *CreateResponse {
	return &CreateResponse{}
}
