// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// Payment represents class Payment
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_Payment
type Payment struct {
	ID            *string       `json:"id,omitempty"`
	PaymentOutput *Output       `json:"paymentOutput,omitempty"`
	Status        *string       `json:"status,omitempty"`
	StatusOutput  *StatusOutput `json:"statusOutput,omitempty"`
}

// NewPayment constructs a new Payment
func NewPayment() *Payment {
	return &Payment{}
}
