// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package token

// MandateApproval represents class MandateApproval
type MandateApproval struct {
	MandateSignatureDate  *string `json:"mandateSignatureDate,omitempty"`
	MandateSignaturePlace *string `json:"mandateSignaturePlace,omitempty"`
	MandateSigned         *bool   `json:"mandateSigned,omitempty"`
}

// NewMandateApproval constructs a new MandateApproval
func NewMandateApproval() *MandateApproval {
	return &MandateApproval{}
}
