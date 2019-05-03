// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// SdkDataOutput represents class SdkDataOutput
type SdkDataOutput struct {
	SdkTransactionID *string `json:"sdkTransactionId,omitempty"`
}

// NewSdkDataOutput constructs a new SdkDataOutput
func NewSdkDataOutput() *SdkDataOutput {
	return &SdkDataOutput{}
}
