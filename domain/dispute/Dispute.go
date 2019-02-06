// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package dispute

// Dispute represents class Dispute
type Dispute struct {
	DisputeOutput *Output       `json:"disputeOutput,omitempty"`
	ID            *string       `json:"id,omitempty"`
	PaymentID     *string       `json:"paymentId,omitempty"`
	Status        *string       `json:"status,omitempty"`
	StatusOutput  *StatusOutput `json:"statusOutput,omitempty"`
}

// NewDispute constructs a new Dispute
func NewDispute() *Dispute {
	return &Dispute{}
}
