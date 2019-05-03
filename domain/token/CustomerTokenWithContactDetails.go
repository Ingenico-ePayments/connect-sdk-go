// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package token

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// CustomerTokenWithContactDetails represents class CustomerTokenWithContactDetails
type CustomerTokenWithContactDetails struct {
	BillingAddress      *definitions.Address            `json:"billingAddress,omitempty"`
	CompanyInformation  *definitions.CompanyInformation `json:"companyInformation,omitempty"`
	ContactDetails      *ContactDetailsToken            `json:"contactDetails,omitempty"`
	MerchantCustomerID  *string                         `json:"merchantCustomerId,omitempty"`
	PersonalInformation *PersonalInformationToken       `json:"personalInformation,omitempty"`
	// Deprecated: Use companyInformation.vatNumber instead
	VatNumber           *string                         `json:"vatNumber,omitempty"`
}

// NewCustomerTokenWithContactDetails constructs a new CustomerTokenWithContactDetails
func NewCustomerTokenWithContactDetails() *CustomerTokenWithContactDetails {
	return &CustomerTokenWithContactDetails{}
}
