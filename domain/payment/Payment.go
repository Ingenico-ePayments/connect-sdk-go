// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// Payment represents class Payment
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
