// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payout

// ApproveRequest represents class ApprovePayoutRequest
type ApproveRequest struct {
	DatePayout *string `json:"datePayout,omitempty"`
}

// NewApproveRequest constructs a new ApproveRequest
func NewApproveRequest() *ApproveRequest {
	return &ApproveRequest{}
}
