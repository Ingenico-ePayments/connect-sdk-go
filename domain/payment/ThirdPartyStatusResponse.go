// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// ThirdPartyStatusResponse represents class ThirdPartyStatusResponse
type ThirdPartyStatusResponse struct {
	ThirdPartyStatus *string `json:"thirdPartyStatus,omitempty"`
}

// NewThirdPartyStatusResponse constructs a new ThirdPartyStatusResponse
func NewThirdPartyStatusResponse() *ThirdPartyStatusResponse {
	return &ThirdPartyStatusResponse{}
}
