// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package productgroups

import (
	communicator "github.com/Ingenico-ePayments/connect-sdk-go/communicator"
)

// FindParams represents query parameters for Get payment product groups
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#__merchantId__productgroups_get
type FindParams struct {
	Amount       *int64
	Hide         []string
	IsRecurring  *bool
	CountryCode  *string
	Locale       *string
	CurrencyCode *string
}

// AddHide adds an element to the Hide array.
func (params *FindParams) AddHide(value string) {
	params.Hide = append(params.Hide, value)
	return
}

// ToRequestParameters converts the query to communicator.RequestParams
func (params *FindParams) ToRequestParameters() communicator.RequestParams {
	reqParams := communicator.RequestParams{}

	communicator.AddRequestParameter(&reqParams, "amount", params.Amount)
	communicator.AddRequestParameter(&reqParams, "hide", params.Hide)
	communicator.AddRequestParameter(&reqParams, "isRecurring", params.IsRecurring)
	communicator.AddRequestParameter(&reqParams, "countryCode", params.CountryCode)
	communicator.AddRequestParameter(&reqParams, "locale", params.Locale)
	communicator.AddRequestParameter(&reqParams, "currencyCode", params.CurrencyCode)

	return reqParams
}

// NewFindParams constructs an instance of FindParams
func NewFindParams() *FindParams {
	return &FindParams{}
}
