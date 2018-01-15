// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package product

// DeviceFingerprintRequest represents class DeviceFingerprintRequest
type DeviceFingerprintRequest struct {
	CollectorCallback *string `json:"collectorCallback,omitempty"`
}

// NewDeviceFingerprintRequest constructs a new DeviceFingerprintRequest
func NewDeviceFingerprintRequest() *DeviceFingerprintRequest {
	return &DeviceFingerprintRequest{}
}
