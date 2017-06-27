// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package token

// ApproveRequest represents class ApproveTokenRequest
type ApproveRequest struct {
	MandateSignatureDate  *string `json:"mandateSignatureDate,omitempty"`
	MandateSignaturePlace *string `json:"mandateSignaturePlace,omitempty"`
	MandateSigned         *bool   `json:"mandateSigned,omitempty"`
}

// NewApproveRequest constructs a new ApproveRequest
func NewApproveRequest() *ApproveRequest {
	return &ApproveRequest{}
}
