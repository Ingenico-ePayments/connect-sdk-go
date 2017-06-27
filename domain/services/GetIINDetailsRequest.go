// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package services

// GetIINDetailsRequest represents class GetIINDetailsRequest
type GetIINDetailsRequest struct {
	Bin            *string         `json:"bin,omitempty"`
	PaymentContext *PaymentContext `json:"paymentContext,omitempty"`
}

// NewGetIINDetailsRequest constructs a new GetIINDetailsRequest
func NewGetIINDetailsRequest() *GetIINDetailsRequest {
	return &GetIINDetailsRequest{}
}
