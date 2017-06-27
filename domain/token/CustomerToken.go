// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package token

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// CustomerToken represents class CustomerToken
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
