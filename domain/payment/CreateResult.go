// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// CreateResult represents class CreatePaymentResult
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_CreatePaymentResult
type CreateResult struct {
	CreationOutput *CreationOutput `json:"creationOutput,omitempty"`
	MerchantAction *MerchantAction `json:"merchantAction,omitempty"`
	Payment        *Payment        `json:"payment,omitempty"`
}

// NewCreateResult constructs a new CreateResult
func NewCreateResult() *CreateResult {
	return &CreateResult{}
}
