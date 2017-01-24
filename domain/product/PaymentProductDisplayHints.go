// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package product

// PaymentProductDisplayHints represents class PaymentProductDisplayHints
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentProductDisplayHints
type PaymentProductDisplayHints struct {
	DisplayOrder *int32  `json:"displayOrder,omitempty"`
	Label        *string `json:"label,omitempty"`
	Logo         *string `json:"logo,omitempty"`
}

// NewPaymentProductDisplayHints constructs a new PaymentProductDisplayHints
func NewPaymentProductDisplayHints() *PaymentProductDisplayHints {
	return &PaymentProductDisplayHints{}
}
