// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package file

// HostedFile represents class HostedFile
type HostedFile struct {
	FileName *string `json:"fileName,omitempty"`
	FileSize *string `json:"fileSize,omitempty"`
	FileType *string `json:"fileType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

// NewHostedFile constructs a new HostedFile
func NewHostedFile() *HostedFile {
	return &HostedFile{}
}
