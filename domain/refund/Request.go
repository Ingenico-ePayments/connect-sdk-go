// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package refund

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// Request represents class RefundRequest
type Request struct {
	AmountOfMoney                 *definitions.AmountOfMoney     `json:"amountOfMoney,omitempty"`
	BankRefundMethodSpecificInput *BankRefundMethodSpecificInput `json:"bankRefundMethodSpecificInput,omitempty"`
	Customer                      *Customer                      `json:"customer,omitempty"`
	RefundDate                    *string                        `json:"refundDate,omitempty"`
	RefundReferences              *References                    `json:"refundReferences,omitempty"`
}

// NewRequest constructs a new Request
func NewRequest() *Request {
	return &Request{}
}
