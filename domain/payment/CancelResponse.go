// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// CancelResponse represents class CancelPaymentResponse
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_CancelPaymentResponse
type CancelResponse struct {
	CardPaymentMethodSpecificOutput   *CancelPaymentCardPaymentMethodSpecificOutput   `json:"cardPaymentMethodSpecificOutput,omitempty"`
	MobilePaymentMethodSpecificOutput *CancelPaymentMobilePaymentMethodSpecificOutput `json:"mobilePaymentMethodSpecificOutput,omitempty"`
	Payment                           *Payment                                        `json:"payment,omitempty"`
}

// NewCancelResponse constructs a new CancelResponse
func NewCancelResponse() *CancelResponse {
	return &CancelResponse{}
}
