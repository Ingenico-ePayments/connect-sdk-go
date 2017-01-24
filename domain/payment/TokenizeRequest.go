// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// TokenizeRequest represents class TokenizePaymentRequest
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_TokenizePaymentRequest
type TokenizeRequest struct {
	Alias *string `json:"alias,omitempty"`
}

// NewTokenizeRequest constructs a new TokenizeRequest
func NewTokenizeRequest() *TokenizeRequest {
	return &TokenizeRequest{}
}
