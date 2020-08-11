// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

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
