// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// BankTransferPaymentMethodSpecificInputBase represents class BankTransferPaymentMethodSpecificInputBase
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_BankTransferPaymentMethodSpecificInputBase
type BankTransferPaymentMethodSpecificInputBase struct {
	AdditionalReference *string `json:"additionalReference,omitempty"`
	PaymentProductID    *int32  `json:"paymentProductId,omitempty"`
}

// NewBankTransferPaymentMethodSpecificInputBase constructs a new BankTransferPaymentMethodSpecificInputBase
func NewBankTransferPaymentMethodSpecificInputBase() *BankTransferPaymentMethodSpecificInputBase {
	return &BankTransferPaymentMethodSpecificInputBase{}
}
