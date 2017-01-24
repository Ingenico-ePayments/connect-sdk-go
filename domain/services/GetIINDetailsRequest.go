// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package services

// GetIINDetailsRequest represents class GetIINDetailsRequest
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_GetIINDetailsRequest
type GetIINDetailsRequest struct {
	Bin            *string         `json:"bin,omitempty"`
	PaymentContext *PaymentContext `json:"paymentContext,omitempty"`
}

// NewGetIINDetailsRequest constructs a new GetIINDetailsRequest
func NewGetIINDetailsRequest() *GetIINDetailsRequest {
	return &GetIINDetailsRequest{}
}
