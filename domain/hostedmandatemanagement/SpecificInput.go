// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package hostedmandatemanagement

// SpecificInput represents class HostedMandateManagementSpecificInput
type SpecificInput struct {
	Locale         *string `json:"locale,omitempty"`
	ReturnURL      *string `json:"returnUrl,omitempty"`
	ShowResultPage *bool   `json:"showResultPage,omitempty"`
	Variant        *string `json:"variant,omitempty"`
}

// NewSpecificInput constructs a new SpecificInput
func NewSpecificInput() *SpecificInput {
	return &SpecificInput{}
}
