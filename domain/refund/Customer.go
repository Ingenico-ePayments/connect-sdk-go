// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package refund

import (
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"
)

// Customer represents class RefundCustomer
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_RefundCustomer
type Customer struct {
	Address            *payment.AddressPersonal        `json:"address,omitempty"`
	CompanyInformation *definitions.CompanyInformation `json:"companyInformation,omitempty"`
	ContactDetails     *definitions.ContactDetailsBase `json:"contactDetails,omitempty"`
}

// NewCustomer constructs a new Customer
func NewCustomer() *Customer {
	return &Customer{}
}
