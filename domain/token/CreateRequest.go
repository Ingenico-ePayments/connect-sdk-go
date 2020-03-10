// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package token

// CreateRequest represents class CreateTokenRequest
type CreateRequest struct {
	Card                   *Card                           `json:"card,omitempty"`
	EWallet                *EWallet                        `json:"eWallet,omitempty"`
	EncryptedCustomerInput *string                         `json:"encryptedCustomerInput,omitempty"`
	NonSepaDirectDebit     *NonSepaDirectDebit             `json:"nonSepaDirectDebit,omitempty"`
	PaymentProductID       *int32                          `json:"paymentProductId,omitempty"`
	SepaDirectDebit        *SepaDirectDebitWithoutCreditor `json:"sepaDirectDebit,omitempty"`
}

// NewCreateRequest constructs a new CreateRequest
func NewCreateRequest() *CreateRequest {
	return &CreateRequest{}
}
