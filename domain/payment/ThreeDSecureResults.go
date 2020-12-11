// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// ThreeDSecureResults represents class ThreeDSecureResults
type ThreeDSecureResults struct {
	AcsTransactionID             *string                    `json:"acsTransactionId,omitempty"`
	AppliedExemption             *string                    `json:"appliedExemption,omitempty"`
	AuthenticationAmount         *definitions.AmountOfMoney `json:"authenticationAmount,omitempty"`
	Cavv                         *string                    `json:"cavv,omitempty"`
	DirectoryServerTransactionID *string                    `json:"directoryServerTransactionId,omitempty"`
	Eci                          *string                    `json:"eci,omitempty"`
	ExemptionOutput              *ExemptionOutput           `json:"exemptionOutput,omitempty"`
	SchemeRiskScore              *int32                     `json:"schemeRiskScore,omitempty"`
	SdkData                      *SdkDataOutput             `json:"sdkData,omitempty"`
	ThreeDSecureData             *ThreeDSecureData          `json:"threeDSecureData,omitempty"`
	ThreeDSecureVersion          *string                    `json:"threeDSecureVersion,omitempty"`
	ThreeDServerTransactionID    *string                    `json:"threeDServerTransactionId,omitempty"`
	Xid                          *string                    `json:"xid,omitempty"`
}

// NewThreeDSecureResults constructs a new ThreeDSecureResults
func NewThreeDSecureResults() *ThreeDSecureResults {
	return &ThreeDSecureResults{}
}
