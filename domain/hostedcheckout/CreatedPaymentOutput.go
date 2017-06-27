// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package hostedcheckout

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"

// CreatedPaymentOutput represents class CreatedPaymentOutput
type CreatedPaymentOutput struct {
	DisplayedData             *DisplayedData              `json:"displayedData,omitempty"`
	Payment                   *payment.Payment            `json:"payment,omitempty"`
	PaymentCreationReferences *payment.CreationReferences `json:"paymentCreationReferences,omitempty"`
	PaymentStatusCategory     *string                     `json:"paymentStatusCategory,omitempty"`
	Tokens                    *string                     `json:"tokens,omitempty"`
}

// NewCreatedPaymentOutput constructs a new CreatedPaymentOutput
func NewCreatedPaymentOutput() *CreatedPaymentOutput {
	return &CreatedPaymentOutput{}
}
