// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package token

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// CustomerTokenWithContactDetails represents class CustomerTokenWithContactDetails
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_CustomerTokenWithContactDetails
type CustomerTokenWithContactDetails struct {
	BillingAddress      *definitions.Address            `json:"billingAddress,omitempty"`
	CompanyInformation  *definitions.CompanyInformation `json:"companyInformation,omitempty"`
	ContactDetails      *ContactDetailsToken            `json:"contactDetails,omitempty"`
	MerchantCustomerID  *string                         `json:"merchantCustomerId,omitempty"`
	PersonalInformation *PersonalInformationToken       `json:"personalInformation,omitempty"`
	VatNumber           *string                         `json:"vatNumber,omitempty"`
}

// NewCustomerTokenWithContactDetails constructs a new CustomerTokenWithContactDetails
func NewCustomerTokenWithContactDetails() *CustomerTokenWithContactDetails {
	return &CustomerTokenWithContactDetails{}
}
