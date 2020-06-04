// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// Output represents class PaymentOutput
type Output struct {
	AmountOfMoney                              *definitions.AmountOfMoney                     `json:"amountOfMoney,omitempty"`
	AmountPaid                                 *int64                                         `json:"amountPaid,omitempty"`
	AmountReversed                             *int64                                         `json:"amountReversed,omitempty"`
	BankTransferPaymentMethodSpecificOutput    *BankTransferPaymentMethodSpecificOutput       `json:"bankTransferPaymentMethodSpecificOutput,omitempty"`
	CardPaymentMethodSpecificOutput            *CardPaymentMethodSpecificOutput               `json:"cardPaymentMethodSpecificOutput,omitempty"`
	CashPaymentMethodSpecificOutput            *CashPaymentMethodSpecificOutput               `json:"cashPaymentMethodSpecificOutput,omitempty"`
	DirectDebitPaymentMethodSpecificOutput     *NonSepaDirectDebitPaymentMethodSpecificOutput `json:"directDebitPaymentMethodSpecificOutput,omitempty"`
	EInvoicePaymentMethodSpecificOutput        *EInvoicePaymentMethodSpecificOutput           `json:"eInvoicePaymentMethodSpecificOutput,omitempty"`
	InvoicePaymentMethodSpecificOutput         *InvoicePaymentMethodSpecificOutput            `json:"invoicePaymentMethodSpecificOutput,omitempty"`
	MobilePaymentMethodSpecificOutput          *MobilePaymentMethodSpecificOutput             `json:"mobilePaymentMethodSpecificOutput,omitempty"`
	PaymentMethod                              *string                                        `json:"paymentMethod,omitempty"`
	RedirectPaymentMethodSpecificOutput        *RedirectPaymentMethodSpecificOutput           `json:"redirectPaymentMethodSpecificOutput,omitempty"`
	References                                 *References                                    `json:"references,omitempty"`
	ReversalReason                             *string                                        `json:"reversalReason,omitempty"`
	SepaDirectDebitPaymentMethodSpecificOutput *SepaDirectDebitPaymentMethodSpecificOutput    `json:"sepaDirectDebitPaymentMethodSpecificOutput,omitempty"`
}

// NewOutput constructs a new Output
func NewOutput() *Output {
	return &Output{}
}
