// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// ApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput represents class ApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_ApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput
type ApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput struct {
	DateCollect *string `json:"dateCollect,omitempty"`
	Token       *string `json:"token,omitempty"`
}

// NewApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput constructs a new ApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput
func NewApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput() *ApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput {
	return &ApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput{}
}
