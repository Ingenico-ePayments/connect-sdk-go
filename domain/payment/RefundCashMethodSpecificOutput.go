// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// RefundCashMethodSpecificOutput represents class RefundCashMethodSpecificOutput
type RefundCashMethodSpecificOutput struct {
	RefundProductID     *int32 `json:"refundProductId,omitempty"`
	TotalAmountPaid     *int64 `json:"totalAmountPaid,omitempty"`
	TotalAmountRefunded *int64 `json:"totalAmountRefunded,omitempty"`
}

// NewRefundCashMethodSpecificOutput constructs a new RefundCashMethodSpecificOutput
func NewRefundCashMethodSpecificOutput() *RefundCashMethodSpecificOutput {
	return &RefundCashMethodSpecificOutput{}
}
