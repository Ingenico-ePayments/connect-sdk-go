package products

import (
	"testing"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator"
)

func TestDirectoryParamsToRequestParameters(t *testing.T) {
	lParams := &DirectoryParams{}
	paramList := communicator.RequestParams{}

	paramRequestCmp(t, lParams, paramList)

	{
		lParams.CurrencyCode = new(string)
		*lParams.CurrencyCode = "EUR"

		param, _ := communicator.NewRequestParam("currencyCode", "EUR")
		paramList = append(paramList, *param)
	}
	paramRequestCmp(t, lParams, paramList)

	{
		lParams.CountryCode = new(string)
		*lParams.CountryCode = "NL"

		param, _ := communicator.NewRequestParam("countryCode", "NL")
		paramList = append(paramList, *param)
	}
	paramRequestCmp(t, lParams, paramList)
}
