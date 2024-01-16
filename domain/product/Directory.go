// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package product

// Directory represents class Directory
type Directory struct {
	Entries *[]DirectoryEntry `json:"entries,omitempty"`
}

// NewDirectory constructs a new Directory
func NewDirectory() *Directory {
	return &Directory{}
}
