// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package dispute

// DisputesResponse represents class DisputesResponse
type DisputesResponse struct {
	Disputes *[]Dispute `json:"disputes,omitempty"`
}

// NewDisputesResponse constructs a new DisputesResponse
func NewDisputesResponse() *DisputesResponse {
	return &DisputesResponse{}
}
