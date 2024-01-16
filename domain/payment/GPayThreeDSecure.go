// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// GPayThreeDSecure represents class GPayThreeDSecure
type GPayThreeDSecure struct {
	ChallengeCanvasSize *string          `json:"challengeCanvasSize,omitempty"`
	ChallengeIndicator  *string          `json:"challengeIndicator,omitempty"`
	ExemptionRequest    *string          `json:"exemptionRequest,omitempty"`
	RedirectionData     *RedirectionData `json:"redirectionData,omitempty"`
	SkipAuthentication  *bool            `json:"skipAuthentication,omitempty"`
}

// NewGPayThreeDSecure constructs a new GPayThreeDSecure
func NewGPayThreeDSecure() *GPayThreeDSecure {
	return &GPayThreeDSecure{}
}
