// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package product

// PaymentProducts represents class PaymentProducts
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentProducts
type PaymentProducts struct {
	PaymentProducts *[]PaymentProduct `json:"paymentProducts,omitempty"`
}

// NewPaymentProducts constructs a new PaymentProducts
func NewPaymentProducts() *PaymentProducts {
	return &PaymentProducts{}
}
