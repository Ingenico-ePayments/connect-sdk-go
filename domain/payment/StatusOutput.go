// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"

// StatusOutput represents class PaymentStatusOutput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentStatusOutput
type StatusOutput struct {
	Errors                   *[]errors.APIError `json:"errors,omitempty"`
	IsAuthorized             *bool              `json:"isAuthorized,omitempty"`
	IsCancellable            *bool              `json:"isCancellable,omitempty"`
	IsRefundable             *bool              `json:"isRefundable,omitempty"`
	StatusCategory           *string            `json:"statusCategory,omitempty"`
	StatusCode               *int32             `json:"statusCode,omitempty"`
	StatusCodeChangeDateTime *string            `json:"statusCodeChangeDateTime,omitempty"`
}

// NewStatusOutput constructs a new StatusOutput
func NewStatusOutput() *StatusOutput {
	return &StatusOutput{}
}
