// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payout

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// CreateRequest represents class CreatePayoutRequest
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_CreatePayoutRequest
type CreateRequest struct {
	AmountOfMoney   *definitions.AmountOfMoney   `json:"amountOfMoney,omitempty"`
	BankAccountBban *definitions.BankAccountBban `json:"bankAccountBban,omitempty"`
	BankAccountIban *definitions.BankAccountIban `json:"bankAccountIban,omitempty"`
	Customer        *Customer                    `json:"customer,omitempty"`
	PayoutDate      *string                      `json:"payoutDate,omitempty"`
	PayoutText      *string                      `json:"payoutText,omitempty"`
	References      *References                  `json:"references,omitempty"`
	SwiftCode       *string                      `json:"swiftCode,omitempty"`
}

// NewCreateRequest constructs a new CreateRequest
func NewCreateRequest() *CreateRequest {
	return &CreateRequest{}
}
