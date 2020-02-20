// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// CashPaymentMethodSpecificOutput represents class CashPaymentMethodSpecificOutput
type CashPaymentMethodSpecificOutput struct {
	FraudResults     *definitions.FraudResults `json:"fraudResults,omitempty"`
	PaymentProductID *int32                    `json:"paymentProductId,omitempty"`
}

// NewCashPaymentMethodSpecificOutput constructs a new CashPaymentMethodSpecificOutput
func NewCashPaymentMethodSpecificOutput() *CashPaymentMethodSpecificOutput {
	return &CashPaymentMethodSpecificOutput{}
}
