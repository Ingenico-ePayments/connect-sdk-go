// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package product

// DirectoryEntry represents class DirectoryEntry
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_DirectoryEntry
type DirectoryEntry struct {
	CountryNames *[]string `json:"countryNames,omitempty"`
	IssuerID     *string   `json:"issuerId,omitempty"`
	IssuerList   *string   `json:"issuerList,omitempty"`
	IssuerName   *string   `json:"issuerName,omitempty"`
}

// NewDirectoryEntry constructs a new DirectoryEntry
func NewDirectoryEntry() *DirectoryEntry {
	return &DirectoryEntry{}
}
