// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package token

// CreateResponse represents class CreateTokenResponse
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_CreateTokenResponse
type CreateResponse struct {
	IsNewToken *bool   `json:"isNewToken,omitempty"`
	Token      *string `json:"token,omitempty"`
}

// NewCreateResponse constructs a new CreateResponse
func NewCreateResponse() *CreateResponse {
	return &CreateResponse{}
}
