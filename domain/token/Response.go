// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package token

// Response represents class TokenResponse
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_TokenResponse
type Response struct {
	Card               *Card               `json:"card,omitempty"`
	EWallet            *EWallet            `json:"eWallet,omitempty"`
	ID                 *string             `json:"id,omitempty"`
	NonSepaDirectDebit *NonSepaDirectDebit `json:"nonSepaDirectDebit,omitempty"`
	PaymentProductID   *int32              `json:"paymentProductId,omitempty"`
	SepaDirectDebit    *SepaDirectDebit    `json:"sepaDirectDebit,omitempty"`
}

// NewResponse constructs a new Response
func NewResponse() *Response {
	return &Response{}
}
