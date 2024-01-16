// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// RefundEInvoiceMethodSpecificOutput represents class RefundEInvoiceMethodSpecificOutput
type RefundEInvoiceMethodSpecificOutput struct {
	RefundProductID     *int32 `json:"refundProductId,omitempty"`
	TotalAmountPaid     *int64 `json:"totalAmountPaid,omitempty"`
	TotalAmountRefunded *int64 `json:"totalAmountRefunded,omitempty"`
}

// NewRefundEInvoiceMethodSpecificOutput constructs a new RefundEInvoiceMethodSpecificOutput
func NewRefundEInvoiceMethodSpecificOutput() *RefundEInvoiceMethodSpecificOutput {
	return &RefundEInvoiceMethodSpecificOutput{}
}
