// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package capture

// Capture represents class Capture
type Capture struct {
	CaptureOutput *Output       `json:"captureOutput,omitempty"`
	ID            *string       `json:"id,omitempty"`
	Status        *string       `json:"status,omitempty"`
	StatusOutput  *StatusOutput `json:"statusOutput,omitempty"`
}

// NewCapture constructs a new Capture
func NewCapture() *Capture {
	return &Capture{}
}
