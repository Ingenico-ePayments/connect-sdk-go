// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// CreateRequest represents class CreatePaymentRequest
type CreateRequest struct {
	BankTransferPaymentMethodSpecificInput    *BankTransferPaymentMethodSpecificInput       `json:"bankTransferPaymentMethodSpecificInput,omitempty"`
	CardPaymentMethodSpecificInput            *CardPaymentMethodSpecificInput               `json:"cardPaymentMethodSpecificInput,omitempty"`
	CashPaymentMethodSpecificInput            *CashPaymentMethodSpecificInput               `json:"cashPaymentMethodSpecificInput,omitempty"`
	DirectDebitPaymentMethodSpecificInput     *NonSepaDirectDebitPaymentMethodSpecificInput `json:"directDebitPaymentMethodSpecificInput,omitempty"`
	EInvoicePaymentMethodSpecificInput        *EInvoicePaymentMethodSpecificInput           `json:"eInvoicePaymentMethodSpecificInput,omitempty"`
	EncryptedCustomerInput                    *string                                       `json:"encryptedCustomerInput,omitempty"`
	FraudFields                               *definitions.FraudFields                      `json:"fraudFields,omitempty"`
	InvoicePaymentMethodSpecificInput         *InvoicePaymentMethodSpecificInput            `json:"invoicePaymentMethodSpecificInput,omitempty"`
	MobilePaymentMethodSpecificInput          *MobilePaymentMethodSpecificInput             `json:"mobilePaymentMethodSpecificInput,omitempty"`
	Order                                     *Order                                        `json:"order,omitempty"`
	RedirectPaymentMethodSpecificInput        *RedirectPaymentMethodSpecificInput           `json:"redirectPaymentMethodSpecificInput,omitempty"`
	SepaDirectDebitPaymentMethodSpecificInput *SepaDirectDebitPaymentMethodSpecificInput    `json:"sepaDirectDebitPaymentMethodSpecificInput,omitempty"`
}

// NewCreateRequest constructs a new CreateRequest
func NewCreateRequest() *CreateRequest {
	return &CreateRequest{}
}
