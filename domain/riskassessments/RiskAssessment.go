// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package riskassessments

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// RiskAssessment represents class RiskAssessment
type RiskAssessment struct {
	FraudFields      *definitions.FraudFields `json:"fraudFields,omitempty"`
	Merchant         *MerchantRiskAssessment  `json:"merchant,omitempty"`
	Order            *OrderRiskAssessment     `json:"order,omitempty"`
	PaymentProductID *int32                   `json:"paymentProductId,omitempty"`
}

// NewRiskAssessment constructs a new RiskAssessment
func NewRiskAssessment() *RiskAssessment {
	return &RiskAssessment{}
}
