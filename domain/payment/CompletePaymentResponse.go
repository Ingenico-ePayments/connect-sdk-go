// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// CompletePaymentResponse represents class CompletePaymentResponse
type CompletePaymentResponse struct {
	CreationOutput *CreationOutput `json:"creationOutput,omitempty"`
	MerchantAction *MerchantAction `json:"merchantAction,omitempty"`
	Payment        *Payment        `json:"payment,omitempty"`
}

// NewCompletePaymentResponse constructs a new CompletePaymentResponse
func NewCompletePaymentResponse() *CompletePaymentResponse {
	return &CompletePaymentResponse{}
}
