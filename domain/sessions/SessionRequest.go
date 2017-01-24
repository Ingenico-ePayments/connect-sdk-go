// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package sessions

// SessionRequest represents class SessionRequest
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_SessionRequest
type SessionRequest struct {
	PaymentProductFilters *PaymentProductFiltersClientSession `json:"paymentProductFilters,omitempty"`
	Tokens                *[]string                           `json:"tokens,omitempty"`
}

// NewSessionRequest constructs a new SessionRequest
func NewSessionRequest() *SessionRequest {
	return &SessionRequest{}
}
