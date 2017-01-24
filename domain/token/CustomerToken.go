// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package token

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// CustomerToken represents class CustomerToken
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_CustomerToken
type CustomerToken struct {
	BillingAddress      *definitions.Address            `json:"billingAddress,omitempty"`
	CompanyInformation  *definitions.CompanyInformation `json:"companyInformation,omitempty"`
	MerchantCustomerID  *string                         `json:"merchantCustomerId,omitempty"`
	PersonalInformation *PersonalInformationToken       `json:"personalInformation,omitempty"`
	VatNumber           *string                         `json:"vatNumber,omitempty"`
}

// NewCustomerToken constructs a new CustomerToken
func NewCustomerToken() *CustomerToken {
	return &CustomerToken{}
}
