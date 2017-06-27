// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// RefundBankMethodSpecificOutput represents class RefundBankMethodSpecificOutput
type RefundBankMethodSpecificOutput struct {
	TotalAmountPaid     *int64 `json:"totalAmountPaid,omitempty"`
	TotalAmountRefunded *int64 `json:"totalAmountRefunded,omitempty"`
}

// NewRefundBankMethodSpecificOutput constructs a new RefundBankMethodSpecificOutput
func NewRefundBankMethodSpecificOutput() *RefundBankMethodSpecificOutput {
	return &RefundBankMethodSpecificOutput{}
}
