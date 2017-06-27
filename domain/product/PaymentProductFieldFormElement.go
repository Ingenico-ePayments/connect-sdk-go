// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

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
