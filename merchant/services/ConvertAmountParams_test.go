package services

import (
	"runtime"
	"testing"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator"
)

func TestToRequestParameters(t *testing.T) {
	lParams := &ConvertAmountParams{}
	paramList := communicator.RequestParams{}

	paramRequestCmp(t, lParams, paramList)

	{
		lParams.Source = new(string)
		*lParams.Source = "EUR"

		param, _ := communicator.NewRequestParam("source", "EUR")
		paramList = append(paramList, *param)
	}
	paramRequestCmp(t, lParams, paramList)

	{
		lParams.Amount = new(int64)
		*lParams.Amount = 1000

		param, _ := communicator.NewRequestParam("amount", "1000")
		paramList = append(paramList, *param)
	}
	paramRequestCmp(t, lParams, paramList)

	{
		lParams.Target = new(string)
		*lParams.Target = "USD"

		param, _ := communicator.NewRequestParam("target", "USD")
		paramList = append(paramList, *param)
	}
	paramRequestCmp(t, lParams, paramList)

	{
		lParams.Source = nil

		paramList = paramList[1:]
	}
	paramRequestCmp(t, lParams, paramList)
}

func paramRequestCmp(t *testing.T, a communicator.ParamRequest, b communicator.RequestParams) {
	params := a.ToRequestParameters()

	if requestParamsCmp(params, b) == false {
		buf := make([]byte, 1<<16)
		runtime.Stack(buf, true)
		t.Fatal("paramRequestCmp failed on equality", params, b, string(buf))
	}
}

func requestParamsCmp(a, b communicator.RequestParams) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
