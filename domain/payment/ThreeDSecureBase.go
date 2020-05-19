// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// ThreeDSecureBase represents class ThreeDSecureBase
type ThreeDSecureBase struct {
	AuthenticationAmount  *definitions.AmountOfMoney `json:"authenticationAmount,omitempty"`
	AuthenticationFlow    *string                    `json:"authenticationFlow,omitempty"`
	ChallengeCanvasSize   *string                    `json:"challengeCanvasSize,omitempty"`
	ChallengeIndicator    *string                    `json:"challengeIndicator,omitempty"`
	ExemptionRequest      *string                    `json:"exemptionRequest,omitempty"`
	PriorThreeDSecureData *ThreeDSecureData          `json:"priorThreeDSecureData,omitempty"`
	SdkData               *SdkDataInput              `json:"sdkData,omitempty"`
	SkipAuthentication    *bool                      `json:"skipAuthentication,omitempty"`
}

// NewThreeDSecureBase constructs a new ThreeDSecureBase
func NewThreeDSecureBase() *ThreeDSecureBase {
	return &ThreeDSecureBase{}
}
