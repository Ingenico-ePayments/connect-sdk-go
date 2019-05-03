// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// MobileThreeDSecureChallengeParameters represents class MobileThreeDSecureChallengeParameters
type MobileThreeDSecureChallengeParameters struct {
	AcsReferenceNumber        *string `json:"acsReferenceNumber,omitempty"`
	AcsSignedContent          *string `json:"acsSignedContent,omitempty"`
	AcsTransactionID          *string `json:"acsTransactionId,omitempty"`
	ThreeDServerTransactionID *string `json:"threeDServerTransactionId,omitempty"`
}

// NewMobileThreeDSecureChallengeParameters constructs a new MobileThreeDSecureChallengeParameters
func NewMobileThreeDSecureChallengeParameters() *MobileThreeDSecureChallengeParameters {
	return &MobileThreeDSecureChallengeParameters{}
}
