// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package dispute

// CreationDetail represents class DisputeCreationDetail
type CreationDetail struct {
	DisputeCreationDate *string `json:"disputeCreationDate,omitempty"`
	DisputeOriginator   *string `json:"disputeOriginator,omitempty"`
	UserName            *string `json:"userName,omitempty"`
}

// NewCreationDetail constructs a new CreationDetail
func NewCreationDetail() *CreationDetail {
	return &CreationDetail{}
}
