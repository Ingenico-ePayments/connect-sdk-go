// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// BankTransferPaymentMethodSpecificInput represents class BankTransferPaymentMethodSpecificInput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_BankTransferPaymentMethodSpecificInput
type BankTransferPaymentMethodSpecificInput struct {
	AdditionalReference *string `json:"additionalReference,omitempty"`
	PaymentProductID    *int32  `json:"paymentProductId,omitempty"`
}

// NewBankTransferPaymentMethodSpecificInput constructs a new BankTransferPaymentMethodSpecificInput
func NewBankTransferPaymentMethodSpecificInput() *BankTransferPaymentMethodSpecificInput {
	return &BankTransferPaymentMethodSpecificInput{}
}
