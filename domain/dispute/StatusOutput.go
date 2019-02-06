// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package dispute

// StatusOutput represents class DisputeStatusOutput
type StatusOutput struct {
	IsCancellable            *bool   `json:"isCancellable,omitempty"`
	StatusCategory           *string `json:"statusCategory,omitempty"`
	StatusCode               *int32  `json:"statusCode,omitempty"`
	StatusCodeChangeDateTime *string `json:"statusCodeChangeDateTime,omitempty"`
}

// NewStatusOutput constructs a new StatusOutput
func NewStatusOutput() *StatusOutput {
	return &StatusOutput{}
}
