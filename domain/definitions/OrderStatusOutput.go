// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package definitions

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"

// OrderStatusOutput represents class OrderStatusOutput
type OrderStatusOutput struct {
	Errors                   *[]errors.APIError `json:"errors,omitempty"`
	IsCancellable            *bool              `json:"isCancellable,omitempty"`
	IsRetriable              *bool              `json:"isRetriable,omitempty"`
	ProviderRawOutput        *[]KeyValuePair    `json:"providerRawOutput,omitempty"`
	StatusCategory           *string            `json:"statusCategory,omitempty"`
	StatusCode               *int32             `json:"statusCode,omitempty"`
	StatusCodeChangeDateTime *string            `json:"statusCodeChangeDateTime,omitempty"`
}

// NewOrderStatusOutput constructs a new OrderStatusOutput
func NewOrderStatusOutput() *OrderStatusOutput {
	return &OrderStatusOutput{}
}
