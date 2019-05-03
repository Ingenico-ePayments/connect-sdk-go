// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// DeviceRenderOptions represents class DeviceRenderOptions
type DeviceRenderOptions struct {
	SdkInterface *string `json:"sdkInterface,omitempty"`
	SdkUIType    *string `json:"sdkUiType,omitempty"`
}

// NewDeviceRenderOptions constructs a new DeviceRenderOptions
func NewDeviceRenderOptions() *DeviceRenderOptions {
	return &DeviceRenderOptions{}
}
