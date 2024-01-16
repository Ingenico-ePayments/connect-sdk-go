// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// ProtectionEligibility represents class ProtectionEligibility
type ProtectionEligibility struct {
	Eligibility *string `json:"eligibility,omitempty"`
	Type        *string `json:"type,omitempty"`
}

// NewProtectionEligibility constructs a new ProtectionEligibility
func NewProtectionEligibility() *ProtectionEligibility {
	return &ProtectionEligibility{}
}
