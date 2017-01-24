// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package token

// MandateApproval represents class MandateApproval
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_MandateApproval
type MandateApproval struct {
	MandateSignatureDate  *string `json:"mandateSignatureDate,omitempty"`
	MandateSignaturePlace *string `json:"mandateSignaturePlace,omitempty"`
	MandateSigned         *bool   `json:"mandateSigned,omitempty"`
}

// NewMandateApproval constructs a new MandateApproval
func NewMandateApproval() *MandateApproval {
	return &MandateApproval{}
}
