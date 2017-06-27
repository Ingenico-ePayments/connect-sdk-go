// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package capture

// StatusOutput represents class CaptureStatusOutput
type StatusOutput struct {
	StatusCode *int32 `json:"statusCode,omitempty"`
}

// NewStatusOutput constructs a new StatusOutput
func NewStatusOutput() *StatusOutput {
	return &StatusOutput{}
}
