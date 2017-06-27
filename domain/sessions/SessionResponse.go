// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package sessions

// SessionResponse represents class SessionResponse
type SessionResponse struct {
	ClientSessionID *string   `json:"clientSessionId,omitempty"`
	CustomerID      *string   `json:"customerId,omitempty"`
	InvalidTokens   *[]string `json:"invalidTokens,omitempty"`
	Region          *string   `json:"region,omitempty"`
}

// NewSessionResponse constructs a new SessionResponse
func NewSessionResponse() *SessionResponse {
	return &SessionResponse{}
}
