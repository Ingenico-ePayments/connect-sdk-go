// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"

// StatusOutput represents class PaymentStatusOutput
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
