// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

// CustomerBase represents class CustomerBase
type CustomerBase struct {
	CompanyInformation *CompanyInformation `json:"companyInformation,omitempty"`
	MerchantCustomerID *string             `json:"merchantCustomerId,omitempty"`
	VatNumber          *string             `json:"vatNumber,omitempty"`
}

// NewCustomerBase constructs a new CustomerBase
func NewCustomerBase() *CustomerBase {
	return &CustomerBase{}
}
