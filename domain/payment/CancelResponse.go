// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// CancelResponse represents class CancelPaymentResponse
type CancelResponse struct {
	CardPaymentMethodSpecificOutput   *CancelPaymentCardPaymentMethodSpecificOutput   `json:"cardPaymentMethodSpecificOutput,omitempty"`
	MobilePaymentMethodSpecificOutput *CancelPaymentMobilePaymentMethodSpecificOutput `json:"mobilePaymentMethodSpecificOutput,omitempty"`
	Payment                           *Payment                                        `json:"payment,omitempty"`
}

// NewCancelResponse constructs a new CancelResponse
func NewCancelResponse() *CancelResponse {
	return &CancelResponse{}
}
