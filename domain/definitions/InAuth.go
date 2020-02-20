// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

// InAuth represents class InAuth
type InAuth struct {
	DeviceCategory       *string `json:"deviceCategory,omitempty"`
	DeviceID             *string `json:"deviceId,omitempty"`
	RiskScore            *string `json:"riskScore,omitempty"`
	TrueIPAddress        *string `json:"trueIpAddress,omitempty"`
	TrueIPAddressCountry *string `json:"trueIpAddressCountry,omitempty"`
}

// NewInAuth constructs a new InAuth
func NewInAuth() *InAuth {
	return &InAuth{}
}
