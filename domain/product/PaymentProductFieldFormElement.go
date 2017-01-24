// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package product

// PaymentProductFieldFormElement represents class PaymentProductFieldFormElement
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentProductFieldFormElement
type PaymentProductFieldFormElement struct {
	Type         *string                `json:"type,omitempty"`
	ValueMapping *[]ValueMappingElement `json:"valueMapping,omitempty"`
}

// NewPaymentProductFieldFormElement constructs a new PaymentProductFieldFormElement
func NewPaymentProductFieldFormElement() *PaymentProductFieldFormElement {
	return &PaymentProductFieldFormElement{}
}
