// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package sessions

// SessionResponse represents class SessionResponse
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_SessionResponse
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
