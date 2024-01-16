// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// DeviceFingerprintDetails represents class DeviceFingerprintDetails
type DeviceFingerprintDetails struct {
	PaymentID                  *string `json:"paymentId,omitempty"`
	RawDeviceFingerprintOutput *string `json:"rawDeviceFingerprintOutput,omitempty"`
}

// NewDeviceFingerprintDetails constructs a new DeviceFingerprintDetails
func NewDeviceFingerprintDetails() *DeviceFingerprintDetails {
	return &DeviceFingerprintDetails{}
}
