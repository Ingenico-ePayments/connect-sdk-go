// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// ApproveRequest represents class ApprovePaymentRequest
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
