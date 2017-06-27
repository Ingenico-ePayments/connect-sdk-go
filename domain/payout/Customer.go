// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payout

import (
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"
)

// Customer represents class PayoutCustomer
type Customer struct {
	Address            *definitions.Address            `json:"address,omitempty"`
	CompanyInformation *definitions.CompanyInformation `json:"companyInformation,omitempty"`
	ContactDetails     *definitions.ContactDetailsBase `json:"contactDetails,omitempty"`
	MerchantCustomerID *string                         `json:"merchantCustomerId,omitempty"`
	Name               *payment.PersonalName           `json:"name,omitempty"`
}

// NewCustomer constructs a new Customer
func NewCustomer() *Customer {
	return &Customer{}
}
