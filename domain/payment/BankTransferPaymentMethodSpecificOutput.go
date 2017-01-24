// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// BankTransferPaymentMethodSpecificOutput represents class BankTransferPaymentMethodSpecificOutput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_BankTransferPaymentMethodSpecificOutput
type BankTransferPaymentMethodSpecificOutput struct {
	PaymentProductID *int32 `json:"paymentProductId,omitempty"`
}

// NewBankTransferPaymentMethodSpecificOutput constructs a new BankTransferPaymentMethodSpecificOutput
func NewBankTransferPaymentMethodSpecificOutput() *BankTransferPaymentMethodSpecificOutput {
	return &BankTransferPaymentMethodSpecificOutput{}
}
