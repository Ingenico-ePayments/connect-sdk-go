// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// RefundBankMethodSpecificOutput represents class RefundBankMethodSpecificOutput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_RefundBankMethodSpecificOutput
type RefundBankMethodSpecificOutput struct {
	TotalAmountPaid     *int64 `json:"totalAmountPaid,omitempty"`
	TotalAmountRefunded *int64 `json:"totalAmountRefunded,omitempty"`
}

// NewRefundBankMethodSpecificOutput constructs a new RefundBankMethodSpecificOutput
func NewRefundBankMethodSpecificOutput() *RefundBankMethodSpecificOutput {
	return &RefundBankMethodSpecificOutput{}
}
