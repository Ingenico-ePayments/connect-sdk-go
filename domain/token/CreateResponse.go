// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package token

// CreateResponse represents class CreateTokenResponse
type CreateResponse struct {
	IsNewToken *bool   `json:"isNewToken,omitempty"`
	Token      *string `json:"token,omitempty"`
}

// NewCreateResponse constructs a new CreateResponse
func NewCreateResponse() *CreateResponse {
	return &CreateResponse{}
}
