// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package services

// GetIINDetailsResponse represents class GetIINDetailsResponse
type GetIINDetailsResponse struct {
	CoBrands           *[]IINDetail `json:"coBrands,omitempty"`
	CountryCode        *string      `json:"countryCode,omitempty"`
	IsAllowedInContext *bool        `json:"isAllowedInContext,omitempty"`
	PaymentProductID   *int32       `json:"paymentProductId,omitempty"`
}

// NewGetIINDetailsResponse constructs a new GetIINDetailsResponse
func NewGetIINDetailsResponse() *GetIINDetailsResponse {
	return &GetIINDetailsResponse{}
}
