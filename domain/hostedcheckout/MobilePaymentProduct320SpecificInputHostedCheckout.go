// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package hostedcheckout

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"

// MobilePaymentProduct320SpecificInputHostedCheckout represents class MobilePaymentProduct320SpecificInputHostedCheckout
type MobilePaymentProduct320SpecificInputHostedCheckout struct {
	MerchantName   *string                   `json:"merchantName,omitempty"`
	MerchantOrigin *string                   `json:"merchantOrigin,omitempty"`
	ThreeDSecure   *payment.GPayThreeDSecure `json:"threeDSecure,omitempty"`
}

// NewMobilePaymentProduct320SpecificInputHostedCheckout constructs a new MobilePaymentProduct320SpecificInputHostedCheckout
func NewMobilePaymentProduct320SpecificInputHostedCheckout() *MobilePaymentProduct320SpecificInputHostedCheckout {
	return &MobilePaymentProduct320SpecificInputHostedCheckout{}
}
