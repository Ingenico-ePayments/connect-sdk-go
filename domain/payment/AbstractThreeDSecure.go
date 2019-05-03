// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// AbstractThreeDSecure represents class AbstractThreeDSecure
type AbstractThreeDSecure struct {
	AuthenticationFlow    *string           `json:"authenticationFlow,omitempty"`
	ChallengeCanvasSize   *string           `json:"challengeCanvasSize,omitempty"`
	ChallengeIndicator    *string           `json:"challengeIndicator,omitempty"`
	PriorThreeDSecureData *ThreeDSecureData `json:"priorThreeDSecureData,omitempty"`
	SdkData               *SdkDataInput     `json:"sdkData,omitempty"`
	SkipAuthentication    *bool             `json:"skipAuthentication,omitempty"`
}

// NewAbstractThreeDSecure constructs a new AbstractThreeDSecure
func NewAbstractThreeDSecure() *AbstractThreeDSecure {
	return &AbstractThreeDSecure{}
}
