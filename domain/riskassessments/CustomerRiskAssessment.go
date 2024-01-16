// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package riskassessments

import (
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"
)

// CustomerRiskAssessment represents class CustomerRiskAssessment
type CustomerRiskAssessment struct {
	Account             *CustomerAccountRiskAssessment     `json:"account,omitempty"`
	AccountType         *string                            `json:"accountType,omitempty"`
	BillingAddress      *definitions.Address               `json:"billingAddress,omitempty"`
	ContactDetails      *ContactDetailsRiskAssessment      `json:"contactDetails,omitempty"`
	Device              *CustomerDeviceRiskAssessment      `json:"device,omitempty"`
	IsPreviousCustomer  *bool                              `json:"isPreviousCustomer,omitempty"`
	Locale              *string                            `json:"locale,omitempty"`
	PersonalInformation *PersonalInformationRiskAssessment `json:"personalInformation,omitempty"`
	// Deprecated: Use Order.shipping.address instead
	ShippingAddress     *payment.AddressPersonal           `json:"shippingAddress,omitempty"`
}

// NewCustomerRiskAssessment constructs a new CustomerRiskAssessment
func NewCustomerRiskAssessment() *CustomerRiskAssessment {
	return &CustomerRiskAssessment{}
}
