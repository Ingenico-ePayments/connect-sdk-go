// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package services

// IINDetail represents class IINDetail
type IINDetail struct {
	IsAllowedInContext *bool  `json:"isAllowedInContext,omitempty"`
	PaymentProductID   *int32 `json:"paymentProductId,omitempty"`
}

// NewIINDetail constructs a new IINDetail
func NewIINDetail() *IINDetail {
	return &IINDetail{}
}
