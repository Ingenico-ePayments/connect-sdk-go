// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package riskassessments

import (
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"
)

// CustomerRiskAssessment represents class CustomerRiskAssessment
type CustomerRiskAssessment struct {
	BillingAddress      *definitions.Address               `json:"billingAddress,omitempty"`
	Locale              *string                            `json:"locale,omitempty"`
	PersonalInformation *PersonalInformationRiskAssessment `json:"personalInformation,omitempty"`
	ShippingAddress     *payment.AddressPersonal           `json:"shippingAddress,omitempty"`
}

// NewCustomerRiskAssessment constructs a new CustomerRiskAssessment
func NewCustomerRiskAssessment() *CustomerRiskAssessment {
	return &CustomerRiskAssessment{}
}
