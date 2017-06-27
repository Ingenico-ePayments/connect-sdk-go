// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package refund

// ApproveRequest represents class ApproveRefundRequest
type ApproveRequest struct {
	Amount *int64 `json:"amount,omitempty"`
}

// NewApproveRequest constructs a new ApproveRequest
func NewApproveRequest() *ApproveRequest {
	return &ApproveRequest{}
}
