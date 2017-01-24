// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// ApprovalResponse represents class PaymentApprovalResponse
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentApprovalResponse
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
