// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// ApproveRequest represents class ApprovePaymentRequest
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_ApprovePaymentRequest
type ApproveRequest struct {
	Amount                                    *int64                                                      `json:"amount,omitempty"`
	DirectDebitPaymentMethodSpecificInput     *ApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput `json:"directDebitPaymentMethodSpecificInput,omitempty"`
	Order                                     *OrderApprovePayment                                        `json:"order,omitempty"`
	SepaDirectDebitPaymentMethodSpecificInput *ApprovePaymentSepaDirectDebitPaymentMethodSpecificInput    `json:"sepaDirectDebitPaymentMethodSpecificInput,omitempty"`
}

// NewApproveRequest constructs a new ApproveRequest
func NewApproveRequest() *ApproveRequest {
	return &ApproveRequest{}
}
