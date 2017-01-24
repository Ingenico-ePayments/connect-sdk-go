// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package definitions

// CompanyInformation represents class CompanyInformation
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_CompanyInformation
type CompanyInformation struct {
	Name *string `json:"name,omitempty"`
}

// NewCompanyInformation constructs a new CompanyInformation
func NewCompanyInformation() *CompanyInformation {
	return &CompanyInformation{}
}
