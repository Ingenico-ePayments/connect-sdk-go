// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// RefundMobileMethodSpecificOutput represents class RefundMobileMethodSpecificOutput
type RefundMobileMethodSpecificOutput struct {
	Network             *string `json:"network,omitempty"`
	TotalAmountPaid     *int64  `json:"totalAmountPaid,omitempty"`
	TotalAmountRefunded *int64  `json:"totalAmountRefunded,omitempty"`
}

// NewRefundMobileMethodSpecificOutput constructs a new RefundMobileMethodSpecificOutput
func NewRefundMobileMethodSpecificOutput() *RefundMobileMethodSpecificOutput {
	return &RefundMobileMethodSpecificOutput{}
}
