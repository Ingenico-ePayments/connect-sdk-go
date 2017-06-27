// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

// FraudFieldsShippingDetails represents class FraudFieldsShippingDetails
type FraudFieldsShippingDetails struct {
	MethodDetails *string `json:"methodDetails,omitempty"`
	MethodSpeed   *int32  `json:"methodSpeed,omitempty"`
	MethodType    *int32  `json:"methodType,omitempty"`
}

// NewFraudFieldsShippingDetails constructs a new FraudFieldsShippingDetails
func NewFraudFieldsShippingDetails() *FraudFieldsShippingDetails {
	return &FraudFieldsShippingDetails{}
}
