// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package products

import "github.com/Ingenico-ePayments/connect-sdk-go/communicator"

// DirectoryParams represents query parameters for Get payment product directory
// Documentation can be found at https://epayments-api.developer-ingenico.com/s2sapi/v1/en_US/go/products/directory.html
type DirectoryParams struct {
	CountryCode  *string
	CurrencyCode *string
}

// ToRequestParameters converts the query to communicator.RequestParams
func (params *DirectoryParams) ToRequestParameters() communicator.RequestParams {
	reqParams := communicator.RequestParams{}

	if params.CountryCode != nil {
		param, _ := communicator.NewRequestParam("countryCode", *params.CountryCode)
		reqParams = append(reqParams, *param)
	}
	if params.CurrencyCode != nil {
		param, _ := communicator.NewRequestParam("currencyCode", *params.CurrencyCode)
		reqParams = append(reqParams, *param)
	}

	return reqParams
}

// NewDirectoryParams constructs an instance of DirectoryParams
func NewDirectoryParams() *DirectoryParams {
	return &DirectoryParams{}
}
