// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package tokens

import "github.com/Ingenico-ePayments/connect-sdk-go/communicator"

// DeleteParams represents query parameters for Delete token
// Documentation can be found at https://epayments-api.developer-ingenico.com/s2sapi/v1/en_US/go/tokens/delete.html
type DeleteParams struct {
	MandateCancelDate *string
}

// ToRequestParameters converts the query to communicator.RequestParams
func (params *DeleteParams) ToRequestParameters() communicator.RequestParams {
	reqParams := communicator.RequestParams{}

	if params.MandateCancelDate != nil {
		param, _ := communicator.NewRequestParam("mandateCancelDate", *params.MandateCancelDate)
		reqParams = append(reqParams, *param)
	}

	return reqParams
}

// NewDeleteParams constructs an instance of DeleteParams
func NewDeleteParams() *DeleteParams {
	return &DeleteParams{}
}
