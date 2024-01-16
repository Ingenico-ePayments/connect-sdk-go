// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package dispute

// DisputesResponse represents class DisputesResponse
type DisputesResponse struct {
	Disputes *[]Dispute `json:"disputes,omitempty"`
}

// NewDisputesResponse constructs a new DisputesResponse
func NewDisputesResponse() *DisputesResponse {
	return &DisputesResponse{}
}
