// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// Customer represents class Customer
type Customer struct {
	Account             *CustomerAccount                `json:"account,omitempty"`
	AccountType         *string                         `json:"accountType,omitempty"`
	BillingAddress      *definitions.Address            `json:"billingAddress,omitempty"`
	CompanyInformation  *definitions.CompanyInformation `json:"companyInformation,omitempty"`
	ContactDetails      *ContactDetails                 `json:"contactDetails,omitempty"`
	Device              *CustomerDevice                 `json:"device,omitempty"`
	FiscalNumber        *string                         `json:"fiscalNumber,omitempty"`
	IsCompany           *bool                           `json:"isCompany,omitempty"`
	IsPreviousCustomer  *bool                           `json:"isPreviousCustomer,omitempty"`
	Locale              *string                         `json:"locale,omitempty"`
	MerchantCustomerID  *string                         `json:"merchantCustomerId,omitempty"`
	PersonalInformation *PersonalInformation            `json:"personalInformation,omitempty"`
	// Deprecated: Use Order.shipping.address instead
	ShippingAddress     *AddressPersonal                `json:"shippingAddress,omitempty"`
	// Deprecated: Use companyInformation.vatNumber instead
	VatNumber           *string                         `json:"vatNumber,omitempty"`
}

// NewCustomer constructs a new Customer
func NewCustomer() *Customer {
	return &Customer{}
}
