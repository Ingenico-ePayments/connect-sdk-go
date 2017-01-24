// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package products

import (
	communicator "github.com/Ingenico-ePayments/connect-sdk-go/communicator"
)

// NetworksParams represents query parameters for Get payment product networks
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#__merchantId__products__paymentProductId__networks_get
type NetworksParams struct {
	CountryCode  *string
	CurrencyCode *string
	Amount       *int64
	IsRecurring  *bool
}

// ToRequestParameters converts the query to communicator.RequestParams
func (params *NetworksParams) ToRequestParameters() communicator.RequestParams {
	reqParams := communicator.RequestParams{}

	communicator.AddRequestParameter(&reqParams, "countryCode", params.CountryCode)
	communicator.AddRequestParameter(&reqParams, "currencyCode", params.CurrencyCode)
	communicator.AddRequestParameter(&reqParams, "amount", params.Amount)
	communicator.AddRequestParameter(&reqParams, "isRecurring", params.IsRecurring)

	return reqParams
}

// NewNetworksParams constructs an instance of NetworksParams
func NewNetworksParams() *NetworksParams {
	return &NetworksParams{}
}
