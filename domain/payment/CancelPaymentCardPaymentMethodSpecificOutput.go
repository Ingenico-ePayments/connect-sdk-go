// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// CancelPaymentCardPaymentMethodSpecificOutput represents class CancelPaymentCardPaymentMethodSpecificOutput
type CancelPaymentCardPaymentMethodSpecificOutput struct {
	VoidResponseID *string `json:"voidResponseId,omitempty"`
}

// NewCancelPaymentCardPaymentMethodSpecificOutput constructs a new CancelPaymentCardPaymentMethodSpecificOutput
func NewCancelPaymentCardPaymentMethodSpecificOutput() *CancelPaymentCardPaymentMethodSpecificOutput {
	return &CancelPaymentCardPaymentMethodSpecificOutput{}
}
