// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// BankTransferPaymentMethodSpecificOutput represents class BankTransferPaymentMethodSpecificOutput
type BankTransferPaymentMethodSpecificOutput struct {
	FraudResults     *definitions.FraudResults `json:"fraudResults,omitempty"`
	PaymentProductID *int32                    `json:"paymentProductId,omitempty"`
}

// NewBankTransferPaymentMethodSpecificOutput constructs a new BankTransferPaymentMethodSpecificOutput
func NewBankTransferPaymentMethodSpecificOutput() *BankTransferPaymentMethodSpecificOutput {
	return &BankTransferPaymentMethodSpecificOutput{}
}
