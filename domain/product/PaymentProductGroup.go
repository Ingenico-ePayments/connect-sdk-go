// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package product

// PaymentProductGroup represents class PaymentProductGroup
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentProductGroup
type PaymentProductGroup struct {
	AccountsOnFile *[]AccountOnFile            `json:"accountsOnFile,omitempty"`
	DisplayHints   *PaymentProductDisplayHints `json:"displayHints,omitempty"`
	Fields         *[]PaymentProductField      `json:"fields,omitempty"`
	ID             *string                     `json:"id,omitempty"`
}

// NewPaymentProductGroup constructs a new PaymentProductGroup
func NewPaymentProductGroup() *PaymentProductGroup {
	return &PaymentProductGroup{}
}
