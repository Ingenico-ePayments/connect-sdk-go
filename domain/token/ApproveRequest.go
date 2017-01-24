// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package token

// ApproveRequest represents class ApproveTokenRequest
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_ApproveTokenRequest
type ApproveRequest struct {
	MandateSignatureDate  *string `json:"mandateSignatureDate,omitempty"`
	MandateSignaturePlace *string `json:"mandateSignaturePlace,omitempty"`
	MandateSigned         *bool   `json:"mandateSigned,omitempty"`
}

// NewApproveRequest constructs a new ApproveRequest
func NewApproveRequest() *ApproveRequest {
	return &ApproveRequest{}
}
