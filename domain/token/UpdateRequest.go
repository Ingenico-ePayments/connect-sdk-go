// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package token

// UpdateRequest represents class UpdateTokenRequest
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_UpdateTokenRequest
type UpdateRequest struct {
	Card               *Card                           `json:"card,omitempty"`
	EWallet            *EWallet                        `json:"eWallet,omitempty"`
	NonSepaDirectDebit *NonSepaDirectDebit             `json:"nonSepaDirectDebit,omitempty"`
	PaymentProductID   *int32                          `json:"paymentProductId,omitempty"`
	SepaDirectDebit    *SepaDirectDebitWithoutCreditor `json:"sepaDirectDebit,omitempty"`
}

// NewUpdateRequest constructs a new UpdateRequest
func NewUpdateRequest() *UpdateRequest {
	return &UpdateRequest{}
}
