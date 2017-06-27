// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package capture

// CapturesResponse represents class CapturesResponse
type CapturesResponse struct {
	Captures *[]Capture `json:"captures,omitempty"`
}

// NewCapturesResponse constructs a new CapturesResponse
func NewCapturesResponse() *CapturesResponse {
	return &CapturesResponse{}
}
