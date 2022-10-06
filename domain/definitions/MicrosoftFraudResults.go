// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

// MicrosoftFraudResults represents class MicrosoftFraudResults
type MicrosoftFraudResults struct {
	DeviceCountryCode *string `json:"deviceCountryCode,omitempty"`
	DeviceID          *string `json:"deviceId,omitempty"`
	FraudScore        *int32  `json:"fraudScore,omitempty"`
	TrueIPAddress     *string `json:"trueIpAddress,omitempty"`
	UserDeviceType    *string `json:"userDeviceType,omitempty"`
}

// NewMicrosoftFraudResults constructs a new MicrosoftFraudResults
func NewMicrosoftFraudResults() *MicrosoftFraudResults {
	return &MicrosoftFraudResults{}
}
