// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// Response represents class PaymentResponse
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentResponse
type Response struct {
	ID            *string       `json:"id,omitempty"`
	PaymentOutput *Output       `json:"paymentOutput,omitempty"`
	Status        *string       `json:"status,omitempty"`
	StatusOutput  *StatusOutput `json:"statusOutput,omitempty"`
}

// NewResponse constructs a new Response
func NewResponse() *Response {
	return &Response{}
}
