// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// ExternalCardholderAuthenticationData represents class ExternalCardholderAuthenticationData
type ExternalCardholderAuthenticationData struct {
	AcsTransactionID             *string `json:"acsTransactionId,omitempty"`
	AppliedExemption             *string `json:"appliedExemption,omitempty"`
	Cavv                         *string `json:"cavv,omitempty"`
	CavvAlgorithm                *string `json:"cavvAlgorithm,omitempty"`
	DirectoryServerTransactionID *string `json:"directoryServerTransactionId,omitempty"`
	Eci                          *int32  `json:"eci,omitempty"`
	SchemeRiskScore              *int32  `json:"schemeRiskScore,omitempty"`
	ThreeDSecureVersion          *string `json:"threeDSecureVersion,omitempty"`
	// Deprecated: No replacement
	ThreeDServerTransactionID    *string `json:"threeDServerTransactionId,omitempty"`
	ValidationResult             *string `json:"validationResult,omitempty"`
	Xid                          *string `json:"xid,omitempty"`
}

// NewExternalCardholderAuthenticationData constructs a new ExternalCardholderAuthenticationData
func NewExternalCardholderAuthenticationData() *ExternalCardholderAuthenticationData {
	return &ExternalCardholderAuthenticationData{}
}
