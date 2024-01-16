// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// BankTransferPaymentMethodSpecificInputBase represents class BankTransferPaymentMethodSpecificInputBase
type BankTransferPaymentMethodSpecificInputBase struct {
	AdditionalReference *string `json:"additionalReference,omitempty"`
	PaymentProductID    *int32  `json:"paymentProductId,omitempty"`
}

// NewBankTransferPaymentMethodSpecificInputBase constructs a new BankTransferPaymentMethodSpecificInputBase
func NewBankTransferPaymentMethodSpecificInputBase() *BankTransferPaymentMethodSpecificInputBase {
	return &BankTransferPaymentMethodSpecificInputBase{}
}
