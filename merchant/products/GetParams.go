// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package products

import (
	"strconv"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator"
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
	if params.ForceBasicFlow != nil {
		param, _ := communicator.NewRequestParam("forceBasicFlow", strconv.FormatBool(*params.ForceBasicFlow))
		reqParams = append(reqParams, *param)
	}

	return reqParams
}

// NewGetParams constructs an instance of GetParams
func NewGetParams() *GetParams {
	return &GetParams{}
}
