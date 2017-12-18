// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// HostedCheckoutSpecificOutput represents class HostedCheckoutSpecificOutput
type HostedCheckoutSpecificOutput struct {
	HostedCheckoutID *string `json:"hostedCheckoutId,omitempty"`
	Variant          *string `json:"variant,omitempty"`
}

// NewHostedCheckoutSpecificOutput constructs a new HostedCheckoutSpecificOutput
func NewHostedCheckoutSpecificOutput() *HostedCheckoutSpecificOutput {
	return &HostedCheckoutSpecificOutput{}
}
