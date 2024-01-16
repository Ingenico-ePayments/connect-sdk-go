// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package capture

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// StatusOutput represents class CaptureStatusOutput
type StatusOutput struct {
	IsRetriable       *bool                       `json:"isRetriable,omitempty"`
	ProviderRawOutput *[]definitions.KeyValuePair `json:"providerRawOutput,omitempty"`
	StatusCode        *int32                      `json:"statusCode,omitempty"`
}

// NewStatusOutput constructs a new StatusOutput
func NewStatusOutput() *StatusOutput {
	return &StatusOutput{}
}
