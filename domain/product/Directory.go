// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package product

// Directory represents class Directory
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_Directory
type Directory struct {
	Entries *[]DirectoryEntry `json:"entries,omitempty"`
}

// NewDirectory constructs a new Directory
func NewDirectory() *Directory {
	return &Directory{}
}
