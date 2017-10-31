// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package product

// PaymentProductFieldDisplayElement represents class PaymentProductFieldDisplayElement
type PaymentProductFieldDisplayElement struct {
	ID    *string `json:"id,omitempty"`
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewPaymentProductFieldDisplayElement constructs a new PaymentProductFieldDisplayElement
func NewPaymentProductFieldDisplayElement() *PaymentProductFieldDisplayElement {
	return &PaymentProductFieldDisplayElement{}
}
