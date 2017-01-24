// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// CreateRequest represents class CreatePaymentRequest
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_CreatePaymentRequest
type CreateRequest struct {
	BankTransferPaymentMethodSpecificInput    *BankTransferPaymentMethodSpecificInput       `json:"bankTransferPaymentMethodSpecificInput,omitempty"`
	CardPaymentMethodSpecificInput            *CardPaymentMethodSpecificInput               `json:"cardPaymentMethodSpecificInput,omitempty"`
	CashPaymentMethodSpecificInput            *CashPaymentMethodSpecificInput               `json:"cashPaymentMethodSpecificInput,omitempty"`
	DirectDebitPaymentMethodSpecificInput     *NonSepaDirectDebitPaymentMethodSpecificInput `json:"directDebitPaymentMethodSpecificInput,omitempty"`
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
