// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package riskassessments

// CustomerDeviceRiskAssessment represents class CustomerDeviceRiskAssessment
type CustomerDeviceRiskAssessment struct {
	DefaultFormFill                *string `json:"defaultFormFill,omitempty"`
	DeviceFingerprintTransactionID *string `json:"deviceFingerprintTransactionId,omitempty"`
}

// NewCustomerDeviceRiskAssessment constructs a new CustomerDeviceRiskAssessment
func NewCustomerDeviceRiskAssessment() *CustomerDeviceRiskAssessment {
	return &CustomerDeviceRiskAssessment{}
}
