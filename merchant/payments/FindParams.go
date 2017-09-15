// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payments

import (
	communicator "github.com/Ingenico-ePayments/connect-sdk-go/communicator"
)

// FindParams represents query parameters for Find payments
// Documentation can be found at https://epayments-api.developer-ingenico.com/s2sapi/v1/en_US/go/payments/find.html
type FindParams struct {
	MerchantReference *string
	MerchantOrderID   *int64
	Offset            *int32
	Limit             *int32
}

// ToRequestParameters converts the query to communicator.RequestParams
func (params *FindParams) ToRequestParameters() communicator.RequestParams {
	reqParams := communicator.RequestParams{}

	communicator.AddRequestParameter(&reqParams, "merchantReference", params.MerchantReference)
	communicator.AddRequestParameter(&reqParams, "merchantOrderId", params.MerchantOrderID)
	communicator.AddRequestParameter(&reqParams, "offset", params.Offset)
	communicator.AddRequestParameter(&reqParams, "limit", params.Limit)

	return reqParams
}

// NewFindParams constructs an instance of FindParams
func NewFindParams() *FindParams {
	return &FindParams{}
}
