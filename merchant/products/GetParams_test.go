package products

import (
	"testing"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator"
)

func TestGetParamsToRequestParameters(t *testing.T) {
	lParams := &GetParams{}
	paramList := communicator.RequestParams{}

	paramRequestCmp(t, lParams, paramList)

	{
		amount := new(int64)
		*amount = 1000
		lParams.Amount = amount

		param, _ := communicator.NewRequestParam("amount", "1000")
		paramList = append(paramList, *param)
	}
	paramRequestCmp(t, lParams, paramList)

	{
		lParams.AddHide("fields")

		param, _ := communicator.NewRequestParam("hide", "fields")
		paramList = append(paramList, *param)
	}
	paramRequestCmp(t, lParams, paramList)

	{
		lParams.AddHide("accountsOnFile")

		param, _ := communicator.NewRequestParam("hide", "accountsOnFile")
		paramList = append(paramList, *param)
	}
	paramRequestCmp(t, lParams, paramList)

	{
		isRecurring := new(bool)
		*isRecurring = true
		lParams.IsRecurring = isRecurring

		param, _ := communicator.NewRequestParam("isRecurring", "true")
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

	{
		lParams.Locale = new(string)
		*lParams.Locale = "nl_NL"

		param, _ := communicator.NewRequestParam("locale", "nl_NL")
		paramList = append(paramList, *param)
	}
	paramRequestCmp(t, lParams, paramList)

	{
		lParams.CurrencyCode = new(string)
		*lParams.CurrencyCode = "EUR"

		param, _ := communicator.NewRequestParam("currencyCode", "EUR")
		paramList = append(paramList, *param)
	}
	paramRequestCmp(t, lParams, paramList)

	{
		lParams.Amount = nil

		paramList = paramList[1:]
	}
	paramRequestCmp(t, lParams, paramList)
}
