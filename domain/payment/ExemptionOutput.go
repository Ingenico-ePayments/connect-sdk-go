// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// ExemptionOutput represents class ExemptionOutput
type ExemptionOutput struct {
	ExemptionRaised          *string `json:"exemptionRaised,omitempty"`
	ExemptionRejectionReason *string `json:"exemptionRejectionReason,omitempty"`
	ExemptionRequest         *string `json:"exemptionRequest,omitempty"`
}

// NewExemptionOutput constructs a new ExemptionOutput
func NewExemptionOutput() *ExemptionOutput {
	return &ExemptionOutput{}
}
