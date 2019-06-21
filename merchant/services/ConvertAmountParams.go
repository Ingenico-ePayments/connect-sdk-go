// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package services

import (
	"strconv"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator"
)

// ConvertAmountParams represents query parameters for Convert amount
// Documentation can be found at https://epayments-api.developer-ingenico.com/s2sapi/v1/en_US/go/services/convertAmount.html
type ConvertAmountParams struct {
	Source *string
	Target *string
	Amount *int64
}

// ToRequestParameters converts the query to communicator.RequestParams
func (params *ConvertAmountParams) ToRequestParameters() communicator.RequestParams {
	reqParams := communicator.RequestParams{}

	if params.Source != nil {
		param, _ := communicator.NewRequestParam("source", *params.Source)
		reqParams = append(reqParams, *param)
	}
	if params.Target != nil {
		param, _ := communicator.NewRequestParam("target", *params.Target)
		reqParams = append(reqParams, *param)
	}
	if params.Amount != nil {
		param, _ := communicator.NewRequestParam("amount", strconv.FormatInt(*params.Amount, 10))
		reqParams = append(reqParams, *param)
	}

	return reqParams
}

// NewConvertAmountParams constructs an instance of ConvertAmountParams
func NewConvertAmountParams() *ConvertAmountParams {
	return &ConvertAmountParams{}
}
