// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// TokenizeRequest represents class TokenizePaymentRequest
type TokenizeRequest struct {
	Alias *string `json:"alias,omitempty"`
}

// NewTokenizeRequest constructs a new TokenizeRequest
func NewTokenizeRequest() *TokenizeRequest {
	return &TokenizeRequest{}
}
