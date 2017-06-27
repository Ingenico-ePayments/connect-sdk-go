// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// RefundMethodSpecificOutput represents class RefundMethodSpecificOutput
type RefundMethodSpecificOutput struct {
	TotalAmountPaid     *int64 `json:"totalAmountPaid,omitempty"`
	TotalAmountRefunded *int64 `json:"totalAmountRefunded,omitempty"`
}

// NewRefundMethodSpecificOutput constructs a new RefundMethodSpecificOutput
func NewRefundMethodSpecificOutput() *RefundMethodSpecificOutput {
	return &RefundMethodSpecificOutput{}
}
