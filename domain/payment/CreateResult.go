// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// CreateResult represents class CreatePaymentResult
type CreateResult struct {
	CreationOutput *CreationOutput `json:"creationOutput,omitempty"`
	MerchantAction *MerchantAction `json:"merchantAction,omitempty"`
	Payment        *Payment        `json:"payment,omitempty"`
}

// NewCreateResult constructs a new CreateResult
func NewCreateResult() *CreateResult {
	return &CreateResult{}
}
