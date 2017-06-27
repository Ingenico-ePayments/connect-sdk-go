// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// ApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput represents class ApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput
type ApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput struct {
	DateCollect *string `json:"dateCollect,omitempty"`
	Token       *string `json:"token,omitempty"`
}

// NewApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput constructs a new ApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput
func NewApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput() *ApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput {
	return &ApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput{}
}
