// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package riskassessments

import (
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"
)

// CustomerRiskAssessment represents class CustomerRiskAssessment
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_CustomerRiskAssessment
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
