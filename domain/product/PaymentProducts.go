// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package product

// PaymentProducts represents class PaymentProducts
type PaymentProducts struct {
	PaymentProducts *[]PaymentProduct `json:"paymentProducts,omitempty"`
}

// NewPaymentProducts constructs a new PaymentProducts
func NewPaymentProducts() *PaymentProducts {
	return &PaymentProducts{}
}
