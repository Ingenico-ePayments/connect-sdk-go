// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// RefundMobileMethodSpecificOutput represents class RefundMobileMethodSpecificOutput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_RefundMobileMethodSpecificOutput
type RefundMobileMethodSpecificOutput struct {
	Network             *string `json:"network,omitempty"`
	TotalAmountPaid     *int64  `json:"totalAmountPaid,omitempty"`
	TotalAmountRefunded *int64  `json:"totalAmountRefunded,omitempty"`
}

// NewRefundMobileMethodSpecificOutput constructs a new RefundMobileMethodSpecificOutput
func NewRefundMobileMethodSpecificOutput() *RefundMobileMethodSpecificOutput {
	return &RefundMobileMethodSpecificOutput{}
}
