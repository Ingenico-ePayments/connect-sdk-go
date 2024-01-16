// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package token

// Response represents class TokenResponse
type Response struct {
	Card               *Card               `json:"card,omitempty"`
	EWallet            *EWallet            `json:"eWallet,omitempty"`
	ID                 *string             `json:"id,omitempty"`
	NonSepaDirectDebit *NonSepaDirectDebit `json:"nonSepaDirectDebit,omitempty"`
	OriginalPaymentID  *string             `json:"originalPaymentId,omitempty"`
	PaymentProductID   *int32              `json:"paymentProductId,omitempty"`
	SepaDirectDebit    *SepaDirectDebit    `json:"sepaDirectDebit,omitempty"`
}

// NewResponse constructs a new Response
func NewResponse() *Response {
	return &Response{}
}
