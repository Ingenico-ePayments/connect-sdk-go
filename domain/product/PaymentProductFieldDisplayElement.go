// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package product

// PaymentProductFieldDisplayElement represents class PaymentProductFieldDisplayElement
type PaymentProductFieldDisplayElement struct {
	ID    *string `json:"id,omitempty"`
	Label *string `json:"label,omitempty"`
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewPaymentProductFieldDisplayElement constructs a new PaymentProductFieldDisplayElement
func NewPaymentProductFieldDisplayElement() *PaymentProductFieldDisplayElement {
	return &PaymentProductFieldDisplayElement{}
}
