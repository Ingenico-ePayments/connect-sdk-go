// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package product

// DeviceFingerprintResponse represents class DeviceFingerprintResponse
type DeviceFingerprintResponse struct {
	DeviceFingerprintTransactionID *string `json:"deviceFingerprintTransactionId,omitempty"`
	HTML                           *string `json:"html,omitempty"`
}

// NewDeviceFingerprintResponse constructs a new DeviceFingerprintResponse
func NewDeviceFingerprintResponse() *DeviceFingerprintResponse {
	return &DeviceFingerprintResponse{}
}
