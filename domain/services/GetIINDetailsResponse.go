// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package services

// GetIINDetailsResponse represents class GetIINDetailsResponse
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_GetIINDetailsResponse
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
