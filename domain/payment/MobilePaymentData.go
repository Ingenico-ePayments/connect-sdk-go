// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// MobilePaymentData represents class MobilePaymentData
type MobilePaymentData struct {
	Dpan       *string `json:"dpan,omitempty"`
	ExpiryDate *string `json:"expiryDate,omitempty"`
}

// NewMobilePaymentData constructs a new MobilePaymentData
func NewMobilePaymentData() *MobilePaymentData {
	return &MobilePaymentData{}
}
