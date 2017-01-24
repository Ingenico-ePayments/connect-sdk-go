// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// RefundMethodSpecificOutput represents class RefundMethodSpecificOutput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_RefundMethodSpecificOutput
type RefundMethodSpecificOutput struct {
	TotalAmountPaid     *int64 `json:"totalAmountPaid,omitempty"`
	TotalAmountRefunded *int64 `json:"totalAmountRefunded,omitempty"`
}

// NewRefundMethodSpecificOutput constructs a new RefundMethodSpecificOutput
func NewRefundMethodSpecificOutput() *RefundMethodSpecificOutput {
	return &RefundMethodSpecificOutput{}
}
