// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// ThirdPartyData represents class ThirdPartyData
type ThirdPartyData struct {
	PaymentProduct863 *Product863ThirdPartyData `json:"paymentProduct863,omitempty"`
}

// NewThirdPartyData constructs a new ThirdPartyData
func NewThirdPartyData() *ThirdPartyData {
	return &ThirdPartyData{}
}
