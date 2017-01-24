// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// ApprovePaymentSepaDirectDebitPaymentMethodSpecificInput represents class ApprovePaymentSepaDirectDebitPaymentMethodSpecificInput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_ApprovePaymentSepaDirectDebitPaymentMethodSpecificInput
type ApprovePaymentSepaDirectDebitPaymentMethodSpecificInput struct {
	DateCollect *string `json:"dateCollect,omitempty"`
	Token       *string `json:"token,omitempty"`
}

// NewApprovePaymentSepaDirectDebitPaymentMethodSpecificInput constructs a new ApprovePaymentSepaDirectDebitPaymentMethodSpecificInput
func NewApprovePaymentSepaDirectDebitPaymentMethodSpecificInput() *ApprovePaymentSepaDirectDebitPaymentMethodSpecificInput {
	return &ApprovePaymentSepaDirectDebitPaymentMethodSpecificInput{}
}
