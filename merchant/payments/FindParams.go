// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payments

import (
	"strconv"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator"
)

// FindParams represents query parameters for Find payments
// Documentation can be found at https://epayments-api.developer-ingenico.com/s2sapi/v1/en_US/go/payments/find.html
type FindParams struct {
	HostedCheckoutID  *string
	MerchantReference *string
	MerchantOrderID   *int64
	Offset            *int32
	Limit             *int32
}

// ToRequestParameters converts the query to communicator.RequestParams
func (params *FindParams) ToRequestParameters() communicator.RequestParams {
	reqParams := communicator.RequestParams{}

	if params.HostedCheckoutID != nil {
		param, _ := communicator.NewRequestParam("hostedCheckoutId", *params.HostedCheckoutID)
		reqParams = append(reqParams, *param)
	}
	if params.MerchantReference != nil {
		param, _ := communicator.NewRequestParam("merchantReference", *params.MerchantReference)
		reqParams = append(reqParams, *param)
	}
	if params.MerchantOrderID != nil {
		param, _ := communicator.NewRequestParam("merchantOrderId", strconv.FormatInt(*params.MerchantOrderID, 10))
		reqParams = append(reqParams, *param)
	}
	if params.Offset != nil {
		param, _ := communicator.NewRequestParam("offset", strconv.FormatInt(int64(*params.Offset), 10))
		reqParams = append(reqParams, *param)
	}
	if params.Limit != nil {
		param, _ := communicator.NewRequestParam("limit", strconv.FormatInt(int64(*params.Limit), 10))
		reqParams = append(reqParams, *param)
	}

	return reqParams
}

// NewFindParams constructs an instance of FindParams
func NewFindParams() *FindParams {
	return &FindParams{}
}
