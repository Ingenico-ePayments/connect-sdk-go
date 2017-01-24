// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// ApprovePaymentPaymentMethodSpecificInput represents class ApprovePaymentPaymentMethodSpecificInput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_ApprovePaymentPaymentMethodSpecificInput
type ApprovePaymentPaymentMethodSpecificInput struct {
	DateCollect *string `json:"dateCollect,omitempty"`
	Token       *string `json:"token,omitempty"`
}

// NewApprovePaymentPaymentMethodSpecificInput constructs a new ApprovePaymentPaymentMethodSpecificInput
func NewApprovePaymentPaymentMethodSpecificInput() *ApprovePaymentPaymentMethodSpecificInput {
	return &ApprovePaymentPaymentMethodSpecificInput{}
}
