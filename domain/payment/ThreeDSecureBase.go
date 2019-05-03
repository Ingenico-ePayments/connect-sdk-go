// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// ThreeDSecureBase represents class ThreeDSecureBase
type ThreeDSecureBase struct {
	AuthenticationFlow    *string           `json:"authenticationFlow,omitempty"`
	ChallengeCanvasSize   *string           `json:"challengeCanvasSize,omitempty"`
	ChallengeIndicator    *string           `json:"challengeIndicator,omitempty"`
	PriorThreeDSecureData *ThreeDSecureData `json:"priorThreeDSecureData,omitempty"`
	SdkData               *SdkDataInput     `json:"sdkData,omitempty"`
	SkipAuthentication    *bool             `json:"skipAuthentication,omitempty"`
}

// NewThreeDSecureBase constructs a new ThreeDSecureBase
func NewThreeDSecureBase() *ThreeDSecureBase {
	return &ThreeDSecureBase{}
}
