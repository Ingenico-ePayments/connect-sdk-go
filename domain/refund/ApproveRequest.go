// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package refund

// ApproveRequest represents class ApproveRefundRequest
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_ApproveRefundRequest
type ApproveRequest struct {
	Amount *int64 `json:"amount,omitempty"`
}

// NewApproveRequest constructs a new ApproveRequest
func NewApproveRequest() *ApproveRequest {
	return &ApproveRequest{}
}
