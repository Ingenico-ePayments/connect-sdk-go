// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package product

// PaymentProductDisplayHints represents class PaymentProductDisplayHints
type PaymentProductDisplayHints struct {
	DisplayOrder *int32  `json:"displayOrder,omitempty"`
	Label        *string `json:"label,omitempty"`
	Logo         *string `json:"logo,omitempty"`
}

// NewPaymentProductDisplayHints constructs a new PaymentProductDisplayHints
func NewPaymentProductDisplayHints() *PaymentProductDisplayHints {
	return &PaymentProductDisplayHints{}
}
