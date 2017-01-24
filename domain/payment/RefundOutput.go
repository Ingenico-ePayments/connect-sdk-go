// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// RefundOutput represents class RefundOutput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_RefundOutput
type RefundOutput struct {
	AmountOfMoney                     *definitions.AmountOfMoney         `json:"amountOfMoney,omitempty"`
	AmountPaid                        *int64                             `json:"amountPaid,omitempty"`
	BankRefundMethodSpecificOutput    *RefundBankMethodSpecificOutput    `json:"bankRefundMethodSpecificOutput,omitempty"`
	CardRefundMethodSpecificOutput    *RefundCardMethodSpecificOutput    `json:"cardRefundMethodSpecificOutput,omitempty"`
	EWalletRefundMethodSpecificOutput *RefundEWalletMethodSpecificOutput `json:"eWalletRefundMethodSpecificOutput,omitempty"`
	MobileRefundMethodSpecificOutput  *RefundMobileMethodSpecificOutput  `json:"mobileRefundMethodSpecificOutput,omitempty"`
	PaymentMethod                     *string                            `json:"paymentMethod,omitempty"`
	References                        *References                        `json:"references,omitempty"`
}

// NewRefundOutput constructs a new RefundOutput
func NewRefundOutput() *RefundOutput {
	return &RefundOutput{}
}
