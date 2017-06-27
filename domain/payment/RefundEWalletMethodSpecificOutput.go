// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// RefundEWalletMethodSpecificOutput represents class RefundEWalletMethodSpecificOutput
type RefundEWalletMethodSpecificOutput struct {
	PaymentProduct840SpecificOutput *RefundPaymentProduct840SpecificOutput `json:"paymentProduct840SpecificOutput,omitempty"`
	TotalAmountPaid                 *int64                                 `json:"totalAmountPaid,omitempty"`
	TotalAmountRefunded             *int64                                 `json:"totalAmountRefunded,omitempty"`
}

// NewRefundEWalletMethodSpecificOutput constructs a new RefundEWalletMethodSpecificOutput
func NewRefundEWalletMethodSpecificOutput() *RefundEWalletMethodSpecificOutput {
	return &RefundEWalletMethodSpecificOutput{}
}
