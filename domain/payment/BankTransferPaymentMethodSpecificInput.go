// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// BankTransferPaymentMethodSpecificInput represents class BankTransferPaymentMethodSpecificInput
type BankTransferPaymentMethodSpecificInput struct {
	AdditionalReference *string `json:"additionalReference,omitempty"`
	PaymentProductID    *int32  `json:"paymentProductId,omitempty"`
}

// NewBankTransferPaymentMethodSpecificInput constructs a new BankTransferPaymentMethodSpecificInput
func NewBankTransferPaymentMethodSpecificInput() *BankTransferPaymentMethodSpecificInput {
	return &BankTransferPaymentMethodSpecificInput{}
}
