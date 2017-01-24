// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// RefundCardMethodSpecificOutput represents class RefundCardMethodSpecificOutput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_RefundCardMethodSpecificOutput
type RefundCardMethodSpecificOutput struct {
	TotalAmountPaid     *int64 `json:"totalAmountPaid,omitempty"`
	TotalAmountRefunded *int64 `json:"totalAmountRefunded,omitempty"`
}

// NewRefundCardMethodSpecificOutput constructs a new RefundCardMethodSpecificOutput
func NewRefundCardMethodSpecificOutput() *RefundCardMethodSpecificOutput {
	return &RefundCardMethodSpecificOutput{}
}
