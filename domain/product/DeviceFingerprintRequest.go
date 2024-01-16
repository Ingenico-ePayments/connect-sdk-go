// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package product

// DeviceFingerprintRequest represents class DeviceFingerprintRequest
type DeviceFingerprintRequest struct {
	CollectorCallback *string `json:"collectorCallback,omitempty"`
}

// NewDeviceFingerprintRequest constructs a new DeviceFingerprintRequest
func NewDeviceFingerprintRequest() *DeviceFingerprintRequest {
	return &DeviceFingerprintRequest{}
}
