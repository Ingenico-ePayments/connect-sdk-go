// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package dispute

// UploadDisputeFileResponse represents class UploadDisputeFileResponse
type UploadDisputeFileResponse struct {
	DisputeID *string `json:"disputeId,omitempty"`
	FileID    *string `json:"fileId,omitempty"`
}

// NewUploadDisputeFileResponse constructs a new UploadDisputeFileResponse
func NewUploadDisputeFileResponse() *UploadDisputeFileResponse {
	return &UploadDisputeFileResponse{}
}
