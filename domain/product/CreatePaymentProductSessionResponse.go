// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package product

// CreatePaymentProductSessionResponse represents class CreatePaymentProductSessionResponse
type CreatePaymentProductSessionResponse struct {
	PaymentProductSession302SpecificOutput *MobilePaymentProductSession302SpecificOutput `json:"paymentProductSession302SpecificOutput,omitempty"`
}

// NewCreatePaymentProductSessionResponse constructs a new CreatePaymentProductSessionResponse
func NewCreatePaymentProductSessionResponse() *CreatePaymentProductSessionResponse {
	return &CreatePaymentProductSessionResponse{}
}
