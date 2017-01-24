// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// Output represents class PaymentOutput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentOutput
type Output struct {
	AmountOfMoney                              *definitions.AmountOfMoney                     `json:"amountOfMoney,omitempty"`
	AmountPaid                                 *int64                                         `json:"amountPaid,omitempty"`
	BankTransferPaymentMethodSpecificOutput    *BankTransferPaymentMethodSpecificOutput       `json:"bankTransferPaymentMethodSpecificOutput,omitempty"`
	CardPaymentMethodSpecificOutput            *CardPaymentMethodSpecificOutput               `json:"cardPaymentMethodSpecificOutput,omitempty"`
	CashPaymentMethodSpecificOutput            *CashPaymentMethodSpecificOutput               `json:"cashPaymentMethodSpecificOutput,omitempty"`
	DirectDebitPaymentMethodSpecificOutput     *NonSepaDirectDebitPaymentMethodSpecificOutput `json:"directDebitPaymentMethodSpecificOutput,omitempty"`
	InvoicePaymentMethodSpecificOutput         *InvoicePaymentMethodSpecificOutput            `json:"invoicePaymentMethodSpecificOutput,omitempty"`
	MobilePaymentMethodSpecificOutput          *MobilePaymentMethodSpecificOutput             `json:"mobilePaymentMethodSpecificOutput,omitempty"`
	PaymentMethod                              *string                                        `json:"paymentMethod,omitempty"`
	RedirectPaymentMethodSpecificOutput        *RedirectPaymentMethodSpecificOutput           `json:"redirectPaymentMethodSpecificOutput,omitempty"`
	References                                 *References                                    `json:"references,omitempty"`
	SepaDirectDebitPaymentMethodSpecificOutput *SepaDirectDebitPaymentMethodSpecificOutput    `json:"sepaDirectDebitPaymentMethodSpecificOutput,omitempty"`
}

// NewOutput constructs a new Output
func NewOutput() *Output {
	return &Output{}
}
