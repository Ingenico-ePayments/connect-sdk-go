// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package product

// PaymentProductFieldFormElement represents class PaymentProductFieldFormElement
type PaymentProductFieldFormElement struct {
	Type         *string                `json:"type,omitempty"`
	ValueMapping *[]ValueMappingElement `json:"valueMapping,omitempty"`
}

// NewPaymentProductFieldFormElement constructs a new PaymentProductFieldFormElement
func NewPaymentProductFieldFormElement() *PaymentProductFieldFormElement {
	return &PaymentProductFieldFormElement{}
}
