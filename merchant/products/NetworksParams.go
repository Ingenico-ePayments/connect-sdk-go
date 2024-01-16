// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package products

import (
	"strconv"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator"
)

// NetworksParams represents query parameters for Get payment product networks
// Documentation can be found at https://epayments-api.developer-ingenico.com/s2sapi/v1/en_US/go/products/networks.html
type NetworksParams struct {
	CountryCode  *string
	CurrencyCode *string
	Amount       *int64
	IsRecurring  *bool
}

// ToRequestParameters converts the query to communicator.RequestParams
func (params *NetworksParams) ToRequestParameters() communicator.RequestParams {
	reqParams := communicator.RequestParams{}

	if params.CountryCode != nil {
		param, _ := communicator.NewRequestParam("countryCode", *params.CountryCode)
		reqParams = append(reqParams, *param)
	}
	if params.CurrencyCode != nil {
		param, _ := communicator.NewRequestParam("currencyCode", *params.CurrencyCode)
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

	return reqParams
}

// NewNetworksParams constructs an instance of NetworksParams
func NewNetworksParams() *NetworksParams {
	return &NetworksParams{}
}
