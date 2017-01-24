// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package tokens

import (
	communicator "github.com/Ingenico-ePayments/connect-sdk-go/communicator"
)

// DeleteParams represents query parameters for Delete token
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#__merchantId__tokens__tokenId__delete
type DeleteParams struct {
	MandateCancelDate *string
}

// ToRequestParameters converts the query to communicator.RequestParams
func (params *DeleteParams) ToRequestParameters() communicator.RequestParams {
	reqParams := communicator.RequestParams{}

	communicator.AddRequestParameter(&reqParams, "mandateCancelDate", params.MandateCancelDate)

	return reqParams
}

// NewDeleteParams constructs an instance of DeleteParams
func NewDeleteParams() *DeleteParams {
	return &DeleteParams{}
}
