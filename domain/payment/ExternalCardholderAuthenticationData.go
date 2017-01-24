// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// ExternalCardholderAuthenticationData represents class ExternalCardholderAuthenticationData
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_ExternalCardholderAuthenticationData
type ExternalCardholderAuthenticationData struct {
	Cavv             *string `json:"cavv,omitempty"`
	CavvAlgorithm    *string `json:"cavvAlgorithm,omitempty"`
	Eci              *int32  `json:"eci,omitempty"`
	ValidationResult *string `json:"validationResult,omitempty"`
	Xid              *string `json:"xid,omitempty"`
}

// NewExternalCardholderAuthenticationData constructs a new ExternalCardholderAuthenticationData
func NewExternalCardholderAuthenticationData() *ExternalCardholderAuthenticationData {
	return &ExternalCardholderAuthenticationData{}
}
