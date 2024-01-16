// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// RefundBankMethodSpecificOutput represents class RefundBankMethodSpecificOutput
type RefundBankMethodSpecificOutput struct {
	RefundProductID     *int32 `json:"refundProductId,omitempty"`
	TotalAmountPaid     *int64 `json:"totalAmountPaid,omitempty"`
	TotalAmountRefunded *int64 `json:"totalAmountRefunded,omitempty"`
}

// NewRefundBankMethodSpecificOutput constructs a new RefundBankMethodSpecificOutput
func NewRefundBankMethodSpecificOutput() *RefundBankMethodSpecificOutput {
	return &RefundBankMethodSpecificOutput{}
}
