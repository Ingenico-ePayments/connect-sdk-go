// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package productgroups

import (
	communicator "github.com/Ingenico-ePayments/connect-sdk-go/communicator"
)

// FindParams represents query parameters for Get payment product groups
// Documentation can be found at https://epayments-api.developer-ingenico.com/s2sapi/v1/en_US/go/productgroups/find.html
type FindParams struct {
	CountryCode  *string
	CurrencyCode *string
	Locale       *string
	Amount       *int64
	IsRecurring  *bool
	Hide         []string
}

// AddHide adds an element to the Hide array.
func (params *FindParams) AddHide(value string) {
	params.Hide = append(params.Hide, value)
	return
}

// ToRequestParameters converts the query to communicator.RequestParams
func (params *FindParams) ToRequestParameters() communicator.RequestParams {
	reqParams := communicator.RequestParams{}

	communicator.AddRequestParameter(&reqParams, "countryCode", params.CountryCode)
	communicator.AddRequestParameter(&reqParams, "currencyCode", params.CurrencyCode)
	communicator.AddRequestParameter(&reqParams, "locale", params.Locale)
	communicator.AddRequestParameter(&reqParams, "amount", params.Amount)
	communicator.AddRequestParameter(&reqParams, "isRecurring", params.IsRecurring)
	communicator.AddRequestParameter(&reqParams, "hide", params.Hide)

	return reqParams
}

// NewFindParams constructs an instance of FindParams
func NewFindParams() *FindParams {
	return &FindParams{}
}
