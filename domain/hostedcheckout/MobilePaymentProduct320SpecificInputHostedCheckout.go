// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package hostedcheckout

// MobilePaymentProduct320SpecificInputHostedCheckout represents class MobilePaymentProduct320SpecificInputHostedCheckout
type MobilePaymentProduct320SpecificInputHostedCheckout struct {
	MerchantName   *string `json:"merchantName,omitempty"`
	MerchantOrigin *string `json:"merchantOrigin,omitempty"`
}

// NewMobilePaymentProduct320SpecificInputHostedCheckout constructs a new MobilePaymentProduct320SpecificInputHostedCheckout
func NewMobilePaymentProduct320SpecificInputHostedCheckout() *MobilePaymentProduct320SpecificInputHostedCheckout {
	return &MobilePaymentProduct320SpecificInputHostedCheckout{}
}
