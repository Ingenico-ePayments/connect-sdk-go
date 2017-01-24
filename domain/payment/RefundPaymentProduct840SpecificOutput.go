// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// RefundPaymentProduct840SpecificOutput represents class RefundPaymentProduct840SpecificOutput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_RefundPaymentProduct840SpecificOutput
type RefundPaymentProduct840SpecificOutput struct {
	CustomerAccount *RefundPaymentProduct840CustomerAccount `json:"customerAccount,omitempty"`
}

// NewRefundPaymentProduct840SpecificOutput constructs a new RefundPaymentProduct840SpecificOutput
func NewRefundPaymentProduct840SpecificOutput() *RefundPaymentProduct840SpecificOutput {
	return &RefundPaymentProduct840SpecificOutput{}
}
