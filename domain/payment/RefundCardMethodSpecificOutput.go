// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// RefundCardMethodSpecificOutput represents class RefundCardMethodSpecificOutput
type RefundCardMethodSpecificOutput struct {
	AuthorisationCode   *string `json:"authorisationCode,omitempty"`
	RefundProductID     *int32  `json:"refundProductId,omitempty"`
	TotalAmountPaid     *int64  `json:"totalAmountPaid,omitempty"`
	TotalAmountRefunded *int64  `json:"totalAmountRefunded,omitempty"`
}

// NewRefundCardMethodSpecificOutput constructs a new RefundCardMethodSpecificOutput
func NewRefundCardMethodSpecificOutput() *RefundCardMethodSpecificOutput {
	return &RefundCardMethodSpecificOutput{}
}
