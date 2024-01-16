// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package product

// PaymentProductField represents class PaymentProductField
type PaymentProductField struct {
	DataRestrictions *PaymentProductFieldDataRestrictions `json:"dataRestrictions,omitempty"`
	DisplayHints     *PaymentProductFieldDisplayHints     `json:"displayHints,omitempty"`
	ID               *string                              `json:"id,omitempty"`
	Type             *string                              `json:"type,omitempty"`
	UsedForLookup    *bool                                `json:"usedForLookup,omitempty"`
}

// NewPaymentProductField constructs a new PaymentProductField
func NewPaymentProductField() *PaymentProductField {
	return &PaymentProductField{}
}
