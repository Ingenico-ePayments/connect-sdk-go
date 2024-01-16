// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// ApprovalResponse represents class PaymentApprovalResponse
type ApprovalResponse struct {
	CardPaymentMethodSpecificOutput   *ApprovePaymentCardPaymentMethodSpecificOutput   `json:"cardPaymentMethodSpecificOutput,omitempty"`
	MobilePaymentMethodSpecificOutput *ApprovePaymentMobilePaymentMethodSpecificOutput `json:"mobilePaymentMethodSpecificOutput,omitempty"`
	Payment                           *Payment                                         `json:"payment,omitempty"`
	// Deprecated: Use cardPaymentMethodSpecificOutput instead
	PaymentMethodSpecificOutput       *ApprovePaymentCardPaymentMethodSpecificOutput   `json:"paymentMethodSpecificOutput,omitempty"`
}

// NewApprovalResponse constructs a new ApprovalResponse
func NewApprovalResponse() *ApprovalResponse {
	return &ApprovalResponse{}
}
