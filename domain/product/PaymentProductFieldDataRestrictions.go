// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package product

// PaymentProductFieldDataRestrictions represents class PaymentProductFieldDataRestrictions
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentProductFieldDataRestrictions
type PaymentProductFieldDataRestrictions struct {
	IsRequired *bool                          `json:"isRequired,omitempty"`
	Validators *PaymentProductFieldValidators `json:"validators,omitempty"`
}

// NewPaymentProductFieldDataRestrictions constructs a new PaymentProductFieldDataRestrictions
func NewPaymentProductFieldDataRestrictions() *PaymentProductFieldDataRestrictions {
	return &PaymentProductFieldDataRestrictions{}
}
