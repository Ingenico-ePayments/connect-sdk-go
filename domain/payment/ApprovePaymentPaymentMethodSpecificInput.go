// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// ApprovePaymentPaymentMethodSpecificInput represents class ApprovePaymentPaymentMethodSpecificInput
type ApprovePaymentPaymentMethodSpecificInput struct {
	DateCollect *string `json:"dateCollect,omitempty"`
	Token       *string `json:"token,omitempty"`
}

// NewApprovePaymentPaymentMethodSpecificInput constructs a new ApprovePaymentPaymentMethodSpecificInput
func NewApprovePaymentPaymentMethodSpecificInput() *ApprovePaymentPaymentMethodSpecificInput {
	return &ApprovePaymentPaymentMethodSpecificInput{}
}
