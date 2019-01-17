// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package hostedcheckout

// MobilePaymentMethodSpecificInputHostedCheckout represents class MobilePaymentMethodSpecificInputHostedCheckout
type MobilePaymentMethodSpecificInputHostedCheckout struct {
	PaymentProduct320SpecificInput *MobilePaymentProduct320SpecificInputHostedCheckout `json:"paymentProduct320SpecificInput,omitempty"`
	PaymentProductID               *int32                                              `json:"paymentProductId,omitempty"`
}

// NewMobilePaymentMethodSpecificInputHostedCheckout constructs a new MobilePaymentMethodSpecificInputHostedCheckout
func NewMobilePaymentMethodSpecificInputHostedCheckout() *MobilePaymentMethodSpecificInputHostedCheckout {
	return &MobilePaymentMethodSpecificInputHostedCheckout{}
}
