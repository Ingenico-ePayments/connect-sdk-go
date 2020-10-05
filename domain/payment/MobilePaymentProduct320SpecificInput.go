// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// MobilePaymentProduct320SpecificInput represents class MobilePaymentProduct320SpecificInput
type MobilePaymentProduct320SpecificInput struct {
	ThreeDSecure *GPayThreeDSecure `json:"threeDSecure,omitempty"`
}

// NewMobilePaymentProduct320SpecificInput constructs a new MobilePaymentProduct320SpecificInput
func NewMobilePaymentProduct320SpecificInput() *MobilePaymentProduct320SpecificInput {
	return &MobilePaymentProduct320SpecificInput{}
}
