// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// MobilePaymentData represents class MobilePaymentData
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_MobilePaymentData
type MobilePaymentData struct {
	Dpan       *string `json:"dpan,omitempty"`
	ExpiryDate *string `json:"expiryDate,omitempty"`
}

// NewMobilePaymentData constructs a new MobilePaymentData
func NewMobilePaymentData() *MobilePaymentData {
	return &MobilePaymentData{}
}
