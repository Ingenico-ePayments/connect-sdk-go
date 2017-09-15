// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// FindResponse represents class FindPaymentsResponse
type FindResponse struct {
	Limit      *int32     `json:"limit,omitempty"`
	Offset     *int32     `json:"offset,omitempty"`
	Payments   *[]Payment `json:"payments,omitempty"`
	TotalCount *int32     `json:"totalCount,omitempty"`
}

// NewFindResponse constructs a new FindResponse
func NewFindResponse() *FindResponse {
	return &FindResponse{}
}
