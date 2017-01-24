// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// CancelPaymentCardPaymentMethodSpecificOutput represents class CancelPaymentCardPaymentMethodSpecificOutput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_CancelPaymentCardPaymentMethodSpecificOutput
type CancelPaymentCardPaymentMethodSpecificOutput struct {
	VoidResponseID *string `json:"voidResponseId,omitempty"`
}

// NewCancelPaymentCardPaymentMethodSpecificOutput constructs a new CancelPaymentCardPaymentMethodSpecificOutput
func NewCancelPaymentCardPaymentMethodSpecificOutput() *CancelPaymentCardPaymentMethodSpecificOutput {
	return &CancelPaymentCardPaymentMethodSpecificOutput{}
}
