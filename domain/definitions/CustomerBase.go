// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package definitions

// CustomerBase represents class CustomerBase
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_CustomerBase
type CustomerBase struct {
	CompanyInformation *CompanyInformation `json:"companyInformation,omitempty"`
	MerchantCustomerID *string             `json:"merchantCustomerId,omitempty"`
	VatNumber          *string             `json:"vatNumber,omitempty"`
}

// NewCustomerBase constructs a new CustomerBase
func NewCustomerBase() *CustomerBase {
	return &CustomerBase{}
}
