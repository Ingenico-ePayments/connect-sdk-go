// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package products

import (
	"strconv"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator"
)

// FindParams represents query parameters for Get payment products
// Documentation can be found at https://epayments-api.developer-ingenico.com/s2sapi/v1/en_US/go/products/find.html
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

	if params.CountryCode != nil {
		param, _ := communicator.NewRequestParam("countryCode", *params.CountryCode)
		reqParams = append(reqParams, *param)
	}
	if params.CurrencyCode != nil {
		param, _ := communicator.NewRequestParam("currencyCode", *params.CurrencyCode)
		reqParams = append(reqParams, *param)
	}
	if params.Locale != nil {
		param, _ := communicator.NewRequestParam("locale", *params.Locale)
		reqParams = append(reqParams, *param)
	}
	if params.Amount != nil {
		param, _ := communicator.NewRequestParam("amount", strconv.FormatInt(*params.Amount, 10))
		reqParams = append(reqParams, *param)
	}
	if params.IsRecurring != nil {
		param, _ := communicator.NewRequestParam("isRecurring", strconv.FormatBool(*params.IsRecurring))
		reqParams = append(reqParams, *param)
	}
	for _, hideElement := range params.Hide {
		param, _ := communicator.NewRequestParam("hide", hideElement)
		reqParams = append(reqParams, *param)
	}

	return reqParams
}

// NewFindParams constructs an instance of FindParams
func NewFindParams() *FindParams {
	return &FindParams{}
}
