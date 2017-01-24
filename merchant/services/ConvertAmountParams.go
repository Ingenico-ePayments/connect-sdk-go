// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package services

import (
	communicator "github.com/Ingenico-ePayments/connect-sdk-go/communicator"
)

// ConvertAmountParams represents query parameters for Convert amount
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#__merchantId__services_convert_amount_get
type ConvertAmountParams struct {
	Source *string
	Amount *int64
	Target *string
}

// ToRequestParameters converts the query to communicator.RequestParams
func (params *ConvertAmountParams) ToRequestParameters() communicator.RequestParams {
	reqParams := communicator.RequestParams{}

	communicator.AddRequestParameter(&reqParams, "source", params.Source)
	communicator.AddRequestParameter(&reqParams, "amount", params.Amount)
	communicator.AddRequestParameter(&reqParams, "target", params.Target)

	return reqParams
}

// NewConvertAmountParams constructs an instance of ConvertAmountParams
func NewConvertAmountParams() *ConvertAmountParams {
	return &ConvertAmountParams{}
}
