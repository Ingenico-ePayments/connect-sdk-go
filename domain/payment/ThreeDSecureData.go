// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// ThreeDSecureData represents class ThreeDSecureData
type ThreeDSecureData struct {
	AcsTransactionID *string `json:"acsTransactionId,omitempty"`
	Method           *string `json:"method,omitempty"`
	UtcTimestamp     *string `json:"utcTimestamp,omitempty"`
}

// NewThreeDSecureData constructs a new ThreeDSecureData
func NewThreeDSecureData() *ThreeDSecureData {
	return &ThreeDSecureData{}
}
