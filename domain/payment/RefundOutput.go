// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// RefundOutput represents class RefundOutput
type RefundOutput struct {
	AmountOfMoney                      *definitions.AmountOfMoney          `json:"amountOfMoney,omitempty"`
	AmountPaid                         *int64                              `json:"amountPaid,omitempty"`
	BankRefundMethodSpecificOutput     *RefundBankMethodSpecificOutput     `json:"bankRefundMethodSpecificOutput,omitempty"`
	CardRefundMethodSpecificOutput     *RefundCardMethodSpecificOutput     `json:"cardRefundMethodSpecificOutput,omitempty"`
	CashRefundMethodSpecificOutput     *RefundCashMethodSpecificOutput     `json:"cashRefundMethodSpecificOutput,omitempty"`
	EInvoiceRefundMethodSpecificOutput *RefundEInvoiceMethodSpecificOutput `json:"eInvoiceRefundMethodSpecificOutput,omitempty"`
	EWalletRefundMethodSpecificOutput  *RefundEWalletMethodSpecificOutput  `json:"eWalletRefundMethodSpecificOutput,omitempty"`
	MobileRefundMethodSpecificOutput   *RefundMobileMethodSpecificOutput   `json:"mobileRefundMethodSpecificOutput,omitempty"`
	PaymentMethod                      *string                             `json:"paymentMethod,omitempty"`
	References                         *References                         `json:"references,omitempty"`
}

// NewRefundOutput constructs a new RefundOutput
func NewRefundOutput() *RefundOutput {
	return &RefundOutput{}
}
