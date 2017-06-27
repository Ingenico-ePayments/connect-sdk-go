// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package capture

import (
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"
)

// Output represents class CaptureOutput
type Output struct {
	AmountOfMoney                              *definitions.AmountOfMoney                             `json:"amountOfMoney,omitempty"`
	AmountPaid                                 *int64                                                 `json:"amountPaid,omitempty"`
	BankTransferPaymentMethodSpecificOutput    *payment.BankTransferPaymentMethodSpecificOutput       `json:"bankTransferPaymentMethodSpecificOutput,omitempty"`
	CardPaymentMethodSpecificOutput            *payment.CardPaymentMethodSpecificOutput               `json:"cardPaymentMethodSpecificOutput,omitempty"`
	CashPaymentMethodSpecificOutput            *payment.CashPaymentMethodSpecificOutput               `json:"cashPaymentMethodSpecificOutput,omitempty"`
	DirectDebitPaymentMethodSpecificOutput     *payment.NonSepaDirectDebitPaymentMethodSpecificOutput `json:"directDebitPaymentMethodSpecificOutput,omitempty"`
	InvoicePaymentMethodSpecificOutput         *payment.InvoicePaymentMethodSpecificOutput            `json:"invoicePaymentMethodSpecificOutput,omitempty"`
	MobilePaymentMethodSpecificOutput          *payment.MobilePaymentMethodSpecificOutput             `json:"mobilePaymentMethodSpecificOutput,omitempty"`
	PaymentMethod                              *string                                                `json:"paymentMethod,omitempty"`
	RedirectPaymentMethodSpecificOutput        *payment.RedirectPaymentMethodSpecificOutput           `json:"redirectPaymentMethodSpecificOutput,omitempty"`
	References                                 *payment.References                                    `json:"references,omitempty"`
	SepaDirectDebitPaymentMethodSpecificOutput *payment.SepaDirectDebitPaymentMethodSpecificOutput    `json:"sepaDirectDebitPaymentMethodSpecificOutput,omitempty"`
}

// NewOutput constructs a new Output
func NewOutput() *Output {
	return &Output{}
}
