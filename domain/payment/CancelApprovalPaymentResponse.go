// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// CancelApprovalPaymentResponse represents class CancelApprovalPaymentResponse
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_CancelApprovalPaymentResponse
type CancelApprovalPaymentResponse struct {
	Payment *Payment `json:"payment,omitempty"`
}

// NewCancelApprovalPaymentResponse constructs a new CancelApprovalPaymentResponse
func NewCancelApprovalPaymentResponse() *CancelApprovalPaymentResponse {
	return &CancelApprovalPaymentResponse{}
}
