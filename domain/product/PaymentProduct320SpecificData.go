// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package product

// PaymentProduct320SpecificData represents class PaymentProduct320SpecificData
type PaymentProduct320SpecificData struct {
	Gateway  *string   `json:"gateway,omitempty"`
	Networks *[]string `json:"networks,omitempty"`
}

// NewPaymentProduct320SpecificData constructs a new PaymentProduct320SpecificData
func NewPaymentProduct320SpecificData() *PaymentProduct320SpecificData {
	return &PaymentProduct320SpecificData{}
}
