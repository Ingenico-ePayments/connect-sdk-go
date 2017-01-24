// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package product

// PaymentProductField represents class PaymentProductField
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentProductField
type PaymentProductField struct {
	DataRestrictions *PaymentProductFieldDataRestrictions `json:"dataRestrictions,omitempty"`
	DisplayHints     *PaymentProductFieldDisplayHints     `json:"displayHints,omitempty"`
	ID               *string                              `json:"id,omitempty"`
	Type             *string                              `json:"type,omitempty"`
}

// NewPaymentProductField constructs a new PaymentProductField
func NewPaymentProductField() *PaymentProductField {
	return &PaymentProductField{}
}
