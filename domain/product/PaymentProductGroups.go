// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package product

// PaymentProductGroups represents class PaymentProductGroups
type PaymentProductGroups struct {
	PaymentProductGroups *[]PaymentProductGroup `json:"paymentProductGroups,omitempty"`
}

// NewPaymentProductGroups constructs a new PaymentProductGroups
func NewPaymentProductGroups() *PaymentProductGroups {
	return &PaymentProductGroups{}
}
