// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// AbstractBankTransferPaymentMethodSpecificInput represents class AbstractBankTransferPaymentMethodSpecificInput
type AbstractBankTransferPaymentMethodSpecificInput struct {
	AdditionalReference *string `json:"additionalReference,omitempty"`
	PaymentProductID    *int32  `json:"paymentProductId,omitempty"`
}

// NewAbstractBankTransferPaymentMethodSpecificInput constructs a new AbstractBankTransferPaymentMethodSpecificInput
func NewAbstractBankTransferPaymentMethodSpecificInput() *AbstractBankTransferPaymentMethodSpecificInput {
	return &AbstractBankTransferPaymentMethodSpecificInput{}
}
