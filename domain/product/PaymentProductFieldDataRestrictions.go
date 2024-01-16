// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package product

// PaymentProductFieldDataRestrictions represents class PaymentProductFieldDataRestrictions
type PaymentProductFieldDataRestrictions struct {
	IsRequired *bool                          `json:"isRequired,omitempty"`
	Validators *PaymentProductFieldValidators `json:"validators,omitempty"`
}

// NewPaymentProductFieldDataRestrictions constructs a new PaymentProductFieldDataRestrictions
func NewPaymentProductFieldDataRestrictions() *PaymentProductFieldDataRestrictions {
	return &PaymentProductFieldDataRestrictions{}
}
