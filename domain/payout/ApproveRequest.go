// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payout

// ApproveRequest represents class ApprovePayoutRequest
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_ApprovePayoutRequest
type ApproveRequest struct {
	DatePayout *string `json:"datePayout,omitempty"`
}

// NewApproveRequest constructs a new ApproveRequest
func NewApproveRequest() *ApproveRequest {
	return &ApproveRequest{}
}
