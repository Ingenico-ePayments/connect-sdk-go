// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package capture

// Response represents class CaptureResponse
type Response struct {
	CaptureOutput *Output       `json:"captureOutput,omitempty"`
	ID            *string       `json:"id,omitempty"`
	Status        *string       `json:"status,omitempty"`
	StatusOutput  *StatusOutput `json:"statusOutput,omitempty"`
}

// NewResponse constructs a new Response
func NewResponse() *Response {
	return &Response{}
}
