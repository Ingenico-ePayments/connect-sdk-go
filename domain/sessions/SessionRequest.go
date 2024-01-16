// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package sessions

// SessionRequest represents class SessionRequest
type SessionRequest struct {
	PaymentProductFilters *PaymentProductFiltersClientSession `json:"paymentProductFilters,omitempty"`
	Tokens                *[]string                           `json:"tokens,omitempty"`
}

// NewSessionRequest constructs a new SessionRequest
func NewSessionRequest() *SessionRequest {
	return &SessionRequest{}
}
