// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package token

// UpdateRequest represents class UpdateTokenRequest
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
