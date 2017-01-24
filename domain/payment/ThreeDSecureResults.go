// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// ThreeDSecureResults represents class ThreeDSecureResults
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_ThreeDSecureResults
type ThreeDSecureResults struct {
	Cavv *string `json:"cavv,omitempty"`
	Eci  *string `json:"eci,omitempty"`
	Xid  *string `json:"xid,omitempty"`
}

// NewThreeDSecureResults constructs a new ThreeDSecureResults
func NewThreeDSecureResults() *ThreeDSecureResults {
	return &ThreeDSecureResults{}
}
