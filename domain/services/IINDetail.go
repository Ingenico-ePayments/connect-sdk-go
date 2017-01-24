// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package services

// IINDetail represents class IINDetail
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_IINDetail
type IINDetail struct {
	IsAllowedInContext *bool  `json:"isAllowedInContext,omitempty"`
	PaymentProductID   *int32 `json:"paymentProductId,omitempty"`
}

// NewIINDetail constructs a new IINDetail
func NewIINDetail() *IINDetail {
	return &IINDetail{}
}
