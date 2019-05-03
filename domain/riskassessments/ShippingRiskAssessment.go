// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package riskassessments

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"

// ShippingRiskAssessment represents class ShippingRiskAssessment
type ShippingRiskAssessment struct {
	Address        *payment.AddressPersonal `json:"address,omitempty"`
	Comments       *string                  `json:"comments,omitempty"`
	TrackingNumber *string                  `json:"trackingNumber,omitempty"`
}

// NewShippingRiskAssessment constructs a new ShippingRiskAssessment
func NewShippingRiskAssessment() *ShippingRiskAssessment {
	return &ShippingRiskAssessment{}
}
