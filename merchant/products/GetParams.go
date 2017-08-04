// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package products

import (
	communicator "github.com/Ingenico-ePayments/connect-sdk-go/communicator"
)

// GetParams represents query parameters for Get payment product
// Documentation can be found at https://epayments-api.developer-ingenico.com/s2sapi/v1/en_US/go/products/get.html
type GetParams struct {
	CountryCode    *string
	CurrencyCode   *string
	Locale         *string
	Amount         *int64
	IsRecurring    *bool
	Hide           []string
	ForceBasicFlow *bool
}

// AddHide adds an element to the Hide array.
func (params *GetParams) AddHide(value string) {
	params.Hide = append(params.Hide, value)
	return
}

// ToRequestParameters converts the query to communicator.RequestParams
func (params *GetParams) ToRequestParameters() communicator.RequestParams {
	reqParams := communicator.RequestParams{}

	communicator.AddRequestParameter(&reqParams, "countryCode", params.CountryCode)
	communicator.AddRequestParameter(&reqParams, "currencyCode", params.CurrencyCode)
	communicator.AddRequestParameter(&reqParams, "locale", params.Locale)
	communicator.AddRequestParameter(&reqParams, "amount", params.Amount)
	communicator.AddRequestParameter(&reqParams, "isRecurring", params.IsRecurring)
	communicator.AddRequestParameter(&reqParams, "hide", params.Hide)
	communicator.AddRequestParameter(&reqParams, "forceBasicFlow", params.ForceBasicFlow)

	return reqParams
}

// NewGetParams constructs an instance of GetParams
func NewGetParams() *GetParams {
	return &GetParams{}
}
