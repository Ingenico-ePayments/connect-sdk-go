// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

// CompanyInformation represents class CompanyInformation
type CompanyInformation struct {
	Name      *string `json:"name,omitempty"`
	VatNumber *string `json:"vatNumber,omitempty"`
}

// NewCompanyInformation constructs a new CompanyInformation
func NewCompanyInformation() *CompanyInformation {
	return &CompanyInformation{}
}
