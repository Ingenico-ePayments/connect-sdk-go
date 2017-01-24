// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package token

// CreateRequest represents class CreateTokenRequest
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_CreateTokenRequest
type CreateRequest struct {
	Card               *Card                           `json:"card,omitempty"`
	EWallet            *EWallet                        `json:"eWallet,omitempty"`
	NonSepaDirectDebit *NonSepaDirectDebit             `json:"nonSepaDirectDebit,omitempty"`
	PaymentProductID   *int32                          `json:"paymentProductId,omitempty"`
	SepaDirectDebit    *SepaDirectDebitWithoutCreditor `json:"sepaDirectDebit,omitempty"`
}

// NewCreateRequest constructs a new CreateRequest
func NewCreateRequest() *CreateRequest {
	return &CreateRequest{}
}
