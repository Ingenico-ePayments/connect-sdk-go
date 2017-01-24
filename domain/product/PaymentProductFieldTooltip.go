// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package product

// PaymentProductFieldTooltip represents class PaymentProductFieldTooltip
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentProductFieldTooltip
type PaymentProductFieldTooltip struct {
	Image *string `json:"image,omitempty"`
	Label *string `json:"label,omitempty"`
}

// NewPaymentProductFieldTooltip constructs a new PaymentProductFieldTooltip
func NewPaymentProductFieldTooltip() *PaymentProductFieldTooltip {
	return &PaymentProductFieldTooltip{}
}
