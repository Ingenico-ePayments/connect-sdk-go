// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package productgroups

import (
	communicator "github.com/Ingenico-ePayments/connect-sdk-go/communicator"
)

// GetParams represents query parameters for Get payment product group
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#__merchantId__productgroups__paymentProductGroupId__get
type GetParams struct {
	Amount       *int64
	Hide         []string
	IsRecurring  *bool
	CountryCode  *string
	Locale       *string
	CurrencyCode *string
}

// AddHide adds an element to the Hide array.
func (params *GetParams) AddHide(value string) {
	params.Hide = append(params.Hide, value)
	return
}

// ToRequestParameters converts the query to communicator.RequestParams
func (params *GetParams) ToRequestParameters() communicator.RequestParams {
	reqParams := communicator.RequestParams{}

	communicator.AddRequestParameter(&reqParams, "amount", params.Amount)
	communicator.AddRequestParameter(&reqParams, "hide", params.Hide)
	communicator.AddRequestParameter(&reqParams, "isRecurring", params.IsRecurring)
	communicator.AddRequestParameter(&reqParams, "countryCode", params.CountryCode)
	communicator.AddRequestParameter(&reqParams, "locale", params.Locale)
	communicator.AddRequestParameter(&reqParams, "currencyCode", params.CurrencyCode)

	return reqParams
}

// NewGetParams constructs an instance of GetParams
func NewGetParams() *GetParams {
	return &GetParams{}
}
