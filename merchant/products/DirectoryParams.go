// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package products

import (
	communicator "github.com/Ingenico-ePayments/connect-sdk-go/communicator"
)

// DirectoryParams represents query parameters for Get payment product directory
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#__merchantId__products__paymentProductId__directory_get
type DirectoryParams struct {
	CurrencyCode *string
	CountryCode  *string
}

// ToRequestParameters converts the query to communicator.RequestParams
func (params *DirectoryParams) ToRequestParameters() communicator.RequestParams {
	reqParams := communicator.RequestParams{}

	communicator.AddRequestParameter(&reqParams, "currencyCode", params.CurrencyCode)
	communicator.AddRequestParameter(&reqParams, "countryCode", params.CountryCode)

	return reqParams
}

// NewDirectoryParams constructs an instance of DirectoryParams
func NewDirectoryParams() *DirectoryParams {
	return &DirectoryParams{}
}
