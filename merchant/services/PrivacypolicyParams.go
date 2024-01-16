// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package services

import (
	"strconv"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator"
)

// PrivacypolicyParams represents query parameters for Get privacy policy
// Documentation can be found at https://epayments-api.developer-ingenico.com/s2sapi/v1/en_US/go/services/privacypolicy.html
type PrivacypolicyParams struct {
	Locale           *string
	PaymentProductID *int32
}

// ToRequestParameters converts the query to communicator.RequestParams
func (params *PrivacypolicyParams) ToRequestParameters() communicator.RequestParams {
	reqParams := communicator.RequestParams{}

	if params.Locale != nil {
		param, _ := communicator.NewRequestParam("locale", *params.Locale)
		reqParams = append(reqParams, *param)
	}
	if params.PaymentProductID != nil {
		param, _ := communicator.NewRequestParam("paymentProductId", strconv.FormatInt(int64(*params.PaymentProductID), 10))
		reqParams = append(reqParams, *param)
	}

	return reqParams
}

// NewPrivacypolicyParams constructs an instance of PrivacypolicyParams
func NewPrivacypolicyParams() *PrivacypolicyParams {
	return &PrivacypolicyParams{}
}
