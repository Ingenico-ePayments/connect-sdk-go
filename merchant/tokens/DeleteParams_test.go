package tokens

import (
	"testing"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator"
)

func TestDeleteParamsToRequestParameters(t *testing.T) {
	lParams := &DeleteParams{}
	paramList := communicator.RequestParams{}

	paramRequestCmp(t, lParams, paramList)

	{
		lParams.MandateCancelDate = new(string)
		*lParams.MandateCancelDate = "20160610"

		param, _ := communicator.NewRequestParam("mandateCancelDate", "20160610")
		paramList = append(paramList, *param)
	}
	paramRequestCmp(t, lParams, paramList)

	{
		lParams.MandateCancelDate = new(string)
		*lParams.MandateCancelDate = ""

		paramList = communicator.RequestParams{}
	}
	paramRequestCmp(t, lParams, paramList)
}

func paramRequestCmp(t *testing.T, a communicator.ParamRequest, b communicator.RequestParams) {
	params := a.ToRequestParameters()

	if requestParamsCmp(params, b) == false {
		t.Fatal("paramRequestCmp failed on equality")
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
