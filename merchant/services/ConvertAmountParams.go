// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package services

import (
	communicator "github.com/Ingenico-ePayments/connect-sdk-go/communicator"
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

	communicator.AddRequestParameter(&reqParams, "source", params.Source)
	communicator.AddRequestParameter(&reqParams, "target", params.Target)
	communicator.AddRequestParameter(&reqParams, "amount", params.Amount)

	return reqParams
}

// NewConvertAmountParams constructs an instance of ConvertAmountParams
func NewConvertAmountParams() *ConvertAmountParams {
	return &ConvertAmountParams{}
}
