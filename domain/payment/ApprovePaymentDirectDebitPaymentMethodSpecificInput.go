// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// ApprovePaymentDirectDebitPaymentMethodSpecificInput represents class ApprovePaymentDirectDebitPaymentMethodSpecificInput
type ApprovePaymentDirectDebitPaymentMethodSpecificInput struct {
	DateCollect *string `json:"dateCollect,omitempty"`
	Token       *string `json:"token,omitempty"`
}

// NewApprovePaymentDirectDebitPaymentMethodSpecificInput constructs a new ApprovePaymentDirectDebitPaymentMethodSpecificInput
func NewApprovePaymentDirectDebitPaymentMethodSpecificInput() *ApprovePaymentDirectDebitPaymentMethodSpecificInput {
	return &ApprovePaymentDirectDebitPaymentMethodSpecificInput{}
}
