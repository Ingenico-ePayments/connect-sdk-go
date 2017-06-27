// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package hostedcheckout

import (
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"
)

// CreateRequest represents class CreateHostedCheckoutRequest
type CreateRequest struct {
	BankTransferPaymentMethodSpecificInput *payment.BankTransferPaymentMethodSpecificInputBase `json:"bankTransferPaymentMethodSpecificInput,omitempty"`
	CardPaymentMethodSpecificInput         *payment.CardPaymentMethodSpecificInputBase         `json:"cardPaymentMethodSpecificInput,omitempty"`
	CashPaymentMethodSpecificInput         *payment.CashPaymentMethodSpecificInputBase         `json:"cashPaymentMethodSpecificInput,omitempty"`
	FraudFields                            *definitions.FraudFields                            `json:"fraudFields,omitempty"`
	HostedCheckoutSpecificInput            *SpecificInput                                      `json:"hostedCheckoutSpecificInput,omitempty"`
	Order                                  *payment.Order                                      `json:"order,omitempty"`
	RedirectPaymentMethodSpecificInput     *payment.RedirectPaymentMethodSpecificInputBase     `json:"redirectPaymentMethodSpecificInput,omitempty"`
}

// NewCreateRequest constructs a new CreateRequest
func NewCreateRequest() *CreateRequest {
	return &CreateRequest{}
}
