// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// ApprovePaymentSepaDirectDebitPaymentMethodSpecificInput represents class ApprovePaymentSepaDirectDebitPaymentMethodSpecificInput
type ApprovePaymentSepaDirectDebitPaymentMethodSpecificInput struct {
	DateCollect *string `json:"dateCollect,omitempty"`
	Token       *string `json:"token,omitempty"`
}

// NewApprovePaymentSepaDirectDebitPaymentMethodSpecificInput constructs a new ApprovePaymentSepaDirectDebitPaymentMethodSpecificInput
func NewApprovePaymentSepaDirectDebitPaymentMethodSpecificInput() *ApprovePaymentSepaDirectDebitPaymentMethodSpecificInput {
	return &ApprovePaymentSepaDirectDebitPaymentMethodSpecificInput{}
}
