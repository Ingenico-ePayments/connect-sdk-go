// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// CancelApprovalPaymentResponse represents class CancelApprovalPaymentResponse
type CancelApprovalPaymentResponse struct {
	Payment *Payment `json:"payment,omitempty"`
}

// NewCancelApprovalPaymentResponse constructs a new CancelApprovalPaymentResponse
func NewCancelApprovalPaymentResponse() *CancelApprovalPaymentResponse {
	return &CancelApprovalPaymentResponse{}
}
