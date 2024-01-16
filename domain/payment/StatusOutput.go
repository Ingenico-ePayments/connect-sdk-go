// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

import (
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"
)

// StatusOutput represents class PaymentStatusOutput
type StatusOutput struct {
	Errors                   *[]errors.APIError          `json:"errors,omitempty"`
	IsAuthorized             *bool                       `json:"isAuthorized,omitempty"`
	IsCancellable            *bool                       `json:"isCancellable,omitempty"`
	IsRefundable             *bool                       `json:"isRefundable,omitempty"`
	IsRetriable              *bool                       `json:"isRetriable,omitempty"`
	ProviderRawOutput        *[]definitions.KeyValuePair `json:"providerRawOutput,omitempty"`
	StatusCategory           *string                     `json:"statusCategory,omitempty"`
	StatusCode               *int32                      `json:"statusCode,omitempty"`
	StatusCodeChangeDateTime *string                     `json:"statusCodeChangeDateTime,omitempty"`
	ThreeDSecureStatus       *string                     `json:"threeDSecureStatus,omitempty"`
}

// NewStatusOutput constructs a new StatusOutput
func NewStatusOutput() *StatusOutput {
	return &StatusOutput{}
}
