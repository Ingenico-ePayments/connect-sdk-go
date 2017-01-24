// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// Customer represents class Customer
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_Customer
type Customer struct {
	BillingAddress      *definitions.Address            `json:"billingAddress,omitempty"`
	CompanyInformation  *definitions.CompanyInformation `json:"companyInformation,omitempty"`
	ContactDetails      *ContactDetails                 `json:"contactDetails,omitempty"`
	FiscalNumber        *string                         `json:"fiscalNumber,omitempty"`
	Locale              *string                         `json:"locale,omitempty"`
	MerchantCustomerID  *string                         `json:"merchantCustomerId,omitempty"`
	PersonalInformation *PersonalInformation            `json:"personalInformation,omitempty"`
	ShippingAddress     *AddressPersonal                `json:"shippingAddress,omitempty"`
	VatNumber           *string                         `json:"vatNumber,omitempty"`
}

// NewCustomer constructs a new Customer
func NewCustomer() *Customer {
	return &Customer{}
}
