// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// CancelPaymentMobilePaymentMethodSpecificOutput represents class CancelPaymentMobilePaymentMethodSpecificOutput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_CancelPaymentMobilePaymentMethodSpecificOutput
type CancelPaymentMobilePaymentMethodSpecificOutput struct {
	VoidResponseID *string `json:"voidResponseId,omitempty"`
}

// NewCancelPaymentMobilePaymentMethodSpecificOutput constructs a new CancelPaymentMobilePaymentMethodSpecificOutput
func NewCancelPaymentMobilePaymentMethodSpecificOutput() *CancelPaymentMobilePaymentMethodSpecificOutput {
	return &CancelPaymentMobilePaymentMethodSpecificOutput{}
}
