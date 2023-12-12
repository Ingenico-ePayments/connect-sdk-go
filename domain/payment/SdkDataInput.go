// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// SdkDataInput represents class SdkDataInput
type SdkDataInput struct {
	DeviceRenderOptions   *DeviceRenderOptions `json:"deviceRenderOptions,omitempty"`
	SdkAppID              *string              `json:"sdkAppId,omitempty"`
	SdkEncryptedData      *string              `json:"sdkEncryptedData,omitempty"`
	SdkEphemeralPublicKey *string              `json:"sdkEphemeralPublicKey,omitempty"`
	SdkMaxTimeout         *string              `json:"sdkMaxTimeout,omitempty"`
	SdkReferenceNumber    *string              `json:"sdkReferenceNumber,omitempty"`
	SdkTransactionID      *string              `json:"sdkTransactionId,omitempty"`
}

// NewSdkDataInput constructs a new SdkDataInput
func NewSdkDataInput() *SdkDataInput {
	return &SdkDataInput{}
}
