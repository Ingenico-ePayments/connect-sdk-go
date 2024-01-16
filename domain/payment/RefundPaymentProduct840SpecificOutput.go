// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// RefundPaymentProduct840SpecificOutput represents class RefundPaymentProduct840SpecificOutput
type RefundPaymentProduct840SpecificOutput struct {
	CustomerAccount *RefundPaymentProduct840CustomerAccount `json:"customerAccount,omitempty"`
}

// NewRefundPaymentProduct840SpecificOutput constructs a new RefundPaymentProduct840SpecificOutput
func NewRefundPaymentProduct840SpecificOutput() *RefundPaymentProduct840SpecificOutput {
	return &RefundPaymentProduct840SpecificOutput{}
}
