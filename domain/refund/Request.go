// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package refund

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// Request represents class RefundRequest
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_RefundRequest
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
