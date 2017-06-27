// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package products

import (
	communicator "github.com/Ingenico-ePayments/connect-sdk-go/communicator"
)

// DirectoryParams represents query parameters for Get payment product directory
// Documentation can be found at https://epayments-api.developer-ingenico.com/s2sapi/v1/en_US/go/products/directory.html
type DirectoryParams struct {
	CountryCode  *string
	CurrencyCode *string
}

// ToRequestParameters converts the query to communicator.RequestParams
func (params *DirectoryParams) ToRequestParameters() communicator.RequestParams {
	reqParams := communicator.RequestParams{}

	communicator.AddRequestParameter(&reqParams, "countryCode", params.CountryCode)
	communicator.AddRequestParameter(&reqParams, "currencyCode", params.CurrencyCode)

	return reqParams
}

// NewDirectoryParams constructs an instance of DirectoryParams
func NewDirectoryParams() *DirectoryParams {
	return &DirectoryParams{}
}
