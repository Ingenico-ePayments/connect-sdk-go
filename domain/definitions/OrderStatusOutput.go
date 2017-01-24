// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package definitions

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"

// OrderStatusOutput represents class OrderStatusOutput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_OrderStatusOutput
type OrderStatusOutput struct {
	Errors                   *[]errors.APIError `json:"errors,omitempty"`
	IsCancellable            *bool              `json:"isCancellable,omitempty"`
	StatusCategory           *string            `json:"statusCategory,omitempty"`
	StatusCode               *int32             `json:"statusCode,omitempty"`
	StatusCodeChangeDateTime *string            `json:"statusCodeChangeDateTime,omitempty"`
}

// NewOrderStatusOutput constructs a new OrderStatusOutput
func NewOrderStatusOutput() *OrderStatusOutput {
	return &OrderStatusOutput{}
}
